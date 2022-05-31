# Start from golang base image
FROM golang:1.18.2-alpine3.15 as dependencies

ENV GO11MODULE=on
ENV GOPROXY="https://goproxy.io,direct"
ENV GOPRIVATE="https://git.robodev.co/*"
ENV GONOSUMDB="git.robodev.co/*"
ENV GITLAB_USERNAME="go_dependency"
ENV GITLAB_TOKEN="x7ss8fHmSLmJfCMo-5Lz"
ENV GOROOT=/usr/local/go
ENV GOPATH=$HOME/go
ENV GOBIN=$GOPATH/bin
ENV PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
# Install git.
# Git is required for fetching the dependencies.

RUN apk update && apk add --no-cache git make gcc libc-dev protobuf-dev protoc grpc

RUN go install github.com/favadi/protoc-go-inject-tag@v1.3.0
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
RUN go install github.com/golang/protobuf/protoc-gen-go@v1.5.2
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

RUN git config \
--global \
url."https://${GITLAB_USERNAME}:${GITLAB_TOKEN}@git.robodev.co".insteadOf \
"https://git.robodev.co"

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

RUN make pbgen

# Build the Go app
RUN make build
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server cmd/**/*.go

# Start a new stage from scratch
# FROM scratch
FROM alpine

RUN GRPC_HEALTH_PROBE_VERSION=v0.3.1 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

# # Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=dependencies /app/bin/server /app/bin/server
COPY --from=dependencies /app/entrypoint.sh /

RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
