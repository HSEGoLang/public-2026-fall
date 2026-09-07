FROM golang:1.27-alpine

RUN CGO_ENABLED=0 go install golang.org/x/tools/cmd/present@v0.49.0

COPY . /lectures

WORKDIR /lectures

ENTRYPOINT ["/go/bin/present", "-base", "/go/pkg/mod/golang.org/x/tools@v0.49.0/cmd/present", "-http", ":8080", "-orighost", "http://pages.manytask.org/hse/go/public-2026-fall", "-use_playground", "-play=0"]
