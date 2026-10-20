#!/usr/bin/env python3
"""Build a static Go present site for any Pages host (requires Go and wget)."""

import argparse
import json
import os
from pathlib import Path
import re
import shutil
import subprocess
import sys
import tempfile
import time
from urllib.parse import quote, urlsplit
from urllib.request import ProxyHandler, build_opener


PRESENT_VERSION = "v0.49.0"
INSTALL_TIMEOUT = 300
BUILD_TIMEOUT = 180
START_TIMEOUT = 30
REQUEST_TIMEOUT = 2


class BuildError(Exception):
    """The site was not built completely; do not publish its partial output."""


def site_prefix(base_url):
    url = urlsplit(base_url)
    if (url.scheme not in ("http", "https") or not url.netloc or url.username
            or url.password or url.query or url.fragment):
        raise BuildError("--base-url must be an absolute HTTP(S) site URL without query or fragment.")
    return url.path.rstrip("/") + "/"


def run(command, description, timeout, **kwargs):
    try:
        result = subprocess.run(command, capture_output=True, text=True, timeout=timeout, **kwargs)
    except subprocess.TimeoutExpired as error:
        raise BuildError(f"{description} timed out after {timeout} seconds.") from error
    except OSError as error:
        raise BuildError(f"Cannot start {description}: {error}.") from error
    if result.returncode:
        diagnostics = (result.stdout + result.stderr)[-6000:].strip()
        raise BuildError(f"{description} failed (exit {result.returncode}).\n{diagnostics}")


def install_present(directory):
    environment = os.environ.copy()
    environment["GOBIN"] = str(directory)
    print(f"Installing Go present {PRESENT_VERSION}...", flush=True)
    run(["go", "install", f"golang.org/x/tools/cmd/present@{PRESENT_VERSION}"],
        "Installing present", INSTALL_TIMEOUT, env=environment, cwd=directory)
    return directory / "present"


def wait_for_present(process, log_path):
    # present chooses a free port itself and logs the resulting address.
    # Ignore HTTP_PROXY for this local connection, just as wget --no-proxy does.
    opener = build_opener(ProxyHandler({}))
    deadline = time.monotonic() + START_TIMEOUT
    while time.monotonic() < deadline:
        log = log_path.read_text(encoding="utf-8", errors="replace")
        if process.poll() is not None:
            raise BuildError(f"present exited before becoming ready.\n{log[-6000:]}")
        match = re.search(r"Open your web browser and visit (http://127\.0\.0\.1:\d+)", log)
        if match:
            try:
                with opener.open(match[1] + "/", timeout=REQUEST_TIMEOUT) as response:
                    if response.status == 200:
                        return match[1]
            except OSError:
                pass
        time.sleep(0.1)
    raise BuildError(f"present did not become ready in {START_TIMEOUT} seconds.\n{log[-6000:]}")


def stop_present(process):
    if process.poll() is None:
        process.terminate()
        try:
            process.wait(timeout=5)
        except subprocess.TimeoutExpired:
            process.kill()
            process.wait(timeout=5)


def finish_site(output, slides, prefix, origin):
    expected = [Path("index.html"), Path("static/slides.js"), Path("static/styles.css")]
    expected += [Path(str(slide) + ".html") for slide in slides]
    for relative in expected:
        path = output / relative
        if not path.is_file() or path.stat().st_size == 0:
            raise BuildError(f"Generated site is missing a required file: {relative}")
    script = output / "static/slides.js"
    text = script.read_text(encoding="utf-8")
    marker = "var PERMANENT_URL_PREFIX = '/static/';"
    if text.count(marker) != 1:
        raise BuildError("Unexpected present slides.js: cannot set the site's asset prefix.")
    text = text.replace(marker, f"var PERMANENT_URL_PREFIX = {json.dumps(prefix + 'static/')};")
    script.write_text(text, encoding="utf-8")
    for path in output.rglob("*.html"):
        if origin in path.read_text(encoding="utf-8"):
            raise BuildError(f"Generated page still links to the build server: {path.relative_to(output)}")
    (output / ".nojekyll").touch()


def build(source, destination, base_url, present=None):
    prefix = site_prefix(base_url)
    source = Path(source).resolve()
    destination = Path(destination).resolve()
    if not source.is_dir():
        raise BuildError(f"Lecture source directory does not exist: {source}")
    if source == destination or source in destination.parents or destination in source.parents:
        raise BuildError("Source and destination must not contain each other.")
    if destination.exists() and (not destination.is_dir() or any(destination.iterdir())):
        raise BuildError(f"Destination must be absent or an empty directory: {destination}")
    slides = sorted(path.relative_to(source) for path in source.rglob("*.slide") if path.is_file())
    if not slides:
        raise BuildError(f"No *.slide presentations found in {source}")
    if not shutil.which("wget"):
        raise BuildError("wget is required to mirror presentations and their resources.")
    if present is not None:
        binary = shutil.which(str(present))
        if binary is None:
            raise BuildError(f"present executable was not found: {present}")
        present = Path(binary).resolve()
    destination.parent.mkdir(parents=True, exist_ok=True)
    # Build beside the destination, then publish it only after all checks pass.
    with tempfile.TemporaryDirectory(prefix=".build-pages-", dir=destination.parent) as temporary:
        temporary = Path(temporary)
        if present is None:
            present = install_present(temporary)
        output = temporary / "site"
        output.mkdir()
        log_path = temporary / "present.log"
        with log_path.open("wb") as log:
            try:
                process = subprocess.Popen(
                    [str(present), "-http=127.0.0.1:0", "-play=false"], cwd=source,
                    stdout=log, stderr=subprocess.STDOUT,
                )
            except OSError as error:
                raise BuildError(f"Cannot start present: {error}") from error
            try:
                origin = wait_for_present(process, log_path)
                print(f"Rendering {len(slides)} presentations...", flush=True)
                # No default five-level recursion limit: nested examples and assets
                # must be included. styles.css is inserted by slides.js at runtime,
                # so it must be an explicit seed for wget's CSS dependency parser.
                # .play code remains visible with -play=false; static hosting has
                # no /compile or WebSocket backend for the Run buttons.
                urls = [origin + "/", origin + "/static/styles.css"]
                urls += [origin + "/" + quote(slide.as_posix()) for slide in slides]
                run([
                    "wget", "--recursive", "--level=inf", "--no-parent", "--no-host-directories",
                    "--adjust-extension", "--convert-links", "--page-requisites", "--no-proxy",
                    "--execute=robots=off", "--timeout=15", "--tries=2", "--no-verbose",
                    "--directory-prefix", str(output), *urls,
                ], "Mirroring presentations", BUILD_TIMEOUT)
                finish_site(output, slides, prefix, origin)
            except BuildError as error:
                diagnostics = log_path.read_text(encoding="utf-8", errors="replace")[-6000:].strip()
                raise BuildError(f"{error}\npresent log:\n{diagnostics}") from error
            finally:
                stop_present(process)
        if destination.exists():
            destination.rmdir()
        output.rename(destination)
    print(f"Built {len(slides)} presentations in {destination} (site path: {prefix}).")


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--source", default="lectures", help="Directory containing *.slide files")
    parser.add_argument("--dest", required=True, help="Absent or empty output directory")
    parser.add_argument("--base-url", required=True, help="Actual public URL of the deployed site")
    parser.add_argument("--present", help="Use an existing present executable instead of installing it")
    arguments = parser.parse_args()
    try:
        build(arguments.source, arguments.dest, arguments.base_url, arguments.present)
    except (BuildError, OSError) as error:
        print(f"Pages build failed: {error}", file=sys.stderr)
        return 1
    return 0


if __name__ == "__main__":
    sys.exit(main())
