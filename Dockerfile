# Build the manager binary
FROM golang:1.15 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# set Proxy
RUN go env -w GOPROXY=https://goproxy.cn,direct
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY conf/ conf/
COPY controllers/ controllers/
COPY filter/ filter/
COPY routers/ routers/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o run main.go

FROM busybox:latest
WORKDIR /
COPY --from=builder /workspace/run .
COPY conf/app.conf conf/app.conf
ENTRYPOINT ["/run"]