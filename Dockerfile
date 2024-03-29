# Start from golang v1.11 base image
FROM golang:1.11 as builder

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . /go/src/github.com/scalog/scalog-operator/

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/scalog/scalog-operator/

# Download dependencies
RUN set -x && \
    go get github.com/golang/dep/cmd/dep && \
    dep ensure -v

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo \
    -ldflags '-extldflags "-static"' -o scalog-operator \
    github.com/scalog/scalog-operator/cmd/manager/

######## Start a new stage from scratch for data layer #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/src/github.com/scalog/scalog-operator/scalog-operator /app/

WORKDIR /app
