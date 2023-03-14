# Build the manager binary
FROM golang:1.19 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

# Copy the go source
COPY main.go main.go

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o testnetlink main.go
# RUN CGO_ENABLED=1 GOOS=linux go build -o manager -a -ldflags '-linkmode external -extldflags "-static"' main.go

FROM alpine:edge
ARG USER=root
RUN apk -U upgrade && apk add --no-cache \
    nmap \
    libcap \
    sudo \
    bash \
    nmap-scripts && \
    setcap cap_net_raw,cap_net_admin,cap_net_bind_service+eip  /usr/bin/nmap \
    && rm -rf /var/cache/apk/*
WORKDIR /
COPY --from=builder /workspace/testnetlink .
USER $USER:$USER

ENTRYPOINT ["/testnetlink"]
