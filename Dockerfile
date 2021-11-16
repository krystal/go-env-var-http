# syntax = docker/dockerfile:1.0-experimental
# Build the manager binary
FROM golang:1.17 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod ./

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY . .
# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o gcs ./

# Create a final image
FROM alpine:3.13.5

WORKDIR /
RUN addgroup --gid 1000 -S gcs && adduser -S gcs -G gcs --uid 1000

COPY --from=builder /workspace/gcs .

USER gcs:gcs
ENTRYPOINT ["/gcs"]
