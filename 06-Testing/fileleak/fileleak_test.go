//go:build linux

package fileleak_test

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"

	"hse-go/06-Testing/fileleak"
)

type fakeT struct {
	failed  bool
	buffer  bytes.Buffer
	cleanup []func()
}

func (f *fakeT) Errorf(msg string, args ...interface{}) {
	f.failed = true
	fmt.Fprintf(&f.buffer, msg, args...)
}

func (f *fakeT) Cleanup(cb func()) {
	f.cleanup = append([]func(){cb}, f.cleanup...)
}

func (f *fakeT) runCleanup() {
	for _, cb := range f.cleanup {
		cb()
	}
}

func checkLeak(t *testing.T, shouldDetect bool, leaker func()) {
	t.Helper()

	var fake fakeT

	fileleak.VerifyNone(&fake)
	leaker()
	fake.runCleanup()

	require.Equal(t, shouldDetect, fake.failed, "diagnostic: %s", fake.buffer.String())
}

func closeFile(f *os.File) {
	if f != nil {
		_ = f.Close()
	}
}

func TestFileLeak_OpenFile(t *testing.T) {
	var f *os.File

	t.Cleanup(func() { closeFile(f) })

	checkLeak(t, true, func() {
		var err error

		f, err = os.Open("/proc/self/exe")
		require.NoError(t, err)
	})
}

func TestFileLeak_AlwaysOpenFile(t *testing.T) {
	f, err := os.Open("/proc/self/exe")
	require.NoError(t, err)

	t.Cleanup(func() { closeFile(f) })
	checkLeak(t, false, func() {})
}

func TestFileLeak_SamePathNewDescriptor(t *testing.T) {
	first, err := os.Open("/proc/self/exe")
	require.NoError(t, err)

	t.Cleanup(func() { closeFile(first) })

	var second *os.File

	t.Cleanup(func() { closeFile(second) })
	checkLeak(t, true, func() {
		second, err = os.Open("/proc/self/exe")
		require.NoError(t, err)
	})
}

func TestFileLeak_ReopenFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "replacement")
	f, err := os.Open("/proc/self/exe")
	require.NoError(t, err)

	t.Cleanup(func() { closeFile(f) })

	checkLeak(t, true, func() {
		require.NoError(t, f.Close())

		f, err = os.Create(path)
		require.NoError(t, err)
	})
}

func TestFileLeak_ClosedFile(t *testing.T) {
	f, err := os.Open("/proc/self/exe")
	require.NoError(t, err)

	t.Cleanup(func() { closeFile(f) })
	checkLeak(t, false, func() { require.NoError(t, f.Close()) })
}

func TestFileLeak_PipeLeak(t *testing.T) {
	var reader, writer *os.File

	t.Cleanup(func() {
		closeFile(reader)
		closeFile(writer)
	})

	checkLeak(t, true, func() {
		var err error

		reader, writer, err = os.Pipe()
		require.NoError(t, err)
		require.NoError(t, reader.Close())
	})
}

func TestFileLeak_PipeNoLeak(t *testing.T) {
	checkLeak(t, false, func() {
		reader, writer, err := os.Pipe()
		require.NoError(t, err)

		t.Cleanup(func() {
			closeFile(reader)
			closeFile(writer)
		})
		require.NoError(t, reader.Close())
		require.NoError(t, writer.Close())
	})
}

func warmNetwork(t *testing.T) {
	t.Helper()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	require.NoError(t, listener.Close())
}

func TestFileLeak_SocketLeak(t *testing.T) {
	warmNetwork(t)

	var listener net.Listener

	t.Cleanup(func() {
		if listener != nil {
			_ = listener.Close()
		}
	})

	checkLeak(t, true, func() {
		var err error

		listener, err = net.Listen("tcp", "127.0.0.1:0")
		require.NoError(t, err)
	})
}

func TestFileLeak_SocketNoLeak(t *testing.T) {
	warmNetwork(t)
	checkLeak(t, false, func() {
		listener, err := net.Listen("tcp", "127.0.0.1:0")
		require.NoError(t, err)

		t.Cleanup(func() { _ = listener.Close() })
		require.NoError(t, listener.Close())
	})
}

func TestFileLeak_DupLeak(t *testing.T) {
	fd := -1

	t.Cleanup(func() {
		if fd >= 0 {
			_ = syscall.Close(fd)
		}
	})

	checkLeak(t, true, func() {
		var err error

		fd, err = syscall.Dup(1)
		require.NoError(t, err)
	})
}

func TestFileLeak_DupNoLeak(t *testing.T) {
	checkLeak(t, false, func() {
		fd, err := syscall.Dup(1)
		require.NoError(t, err)
		require.NoError(t, syscall.Close(fd))
	})
}

func TestFileLeak_CleanupOrder(t *testing.T) {
	var fake fakeT

	fileleak.VerifyNone(&fake)

	f, err := os.Open("/proc/self/exe")
	require.NoError(t, err)

	t.Cleanup(func() { closeFile(f) })
	fake.Cleanup(func() { closeFile(f) })
	fake.runCleanup()

	require.False(t, fake.failed, "diagnostic: %s", fake.buffer.String())
}
