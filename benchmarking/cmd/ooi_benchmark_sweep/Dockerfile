# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get github.com/amarburg/go-lazyfs

## Having weird dependency resolution problem from within packages...
RUN go get github.com/amarburg/go-lazyfs-testfiles
RUN go get github.com/amarburg/go-lazyfs-testfiles/http_server

RUN go install github.com/amarburg/go-lazyfs/benchmarking/cmd/ooi_benchmark_sweep/

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/ooi_benchmark_sweep

# Document that the service listens on port 8080.
#EXPOSE 8080
