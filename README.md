# GRPC Hello World

Following [this project](https://github.com/grpc/grpc-go/tree/master/examples/helloworld)

## Building and running

```bash
# Build all the project
$ make clean && make

# Run the server
$ ./server

# Run the client
$ ./client [name] [invications:10]
```

## Unix sockets

The brach `unixsocket` contains an implementation that use [unix sockets](https://en.wikipedia.org/wiki/Unix_domain_socket) instead of network (*TCP/IP sockets*). It's is built and ran in the same way that the TCP implementation.

To see the unix sockets opened:

```bash
netstat -a -p --unix
```

## Links

- [https://grpc.io/docs/quickstart/go/](https://grpc.io/docs/quickstart/go/)
- [https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases)
- [https://grpc.io/docs/guides/](https://grpc.io/docs/guides/)
- [https://github.com/jeffwubj/grpc-unixsocket](https://github.com/jeffwubj/grpc-unixsocket)

## How to install protoc environment for Go

From [gRPC quickstart Go article](https://grpc.io/docs/quickstart/go/#prerequisites)

### Install the protoc compiler

```bash
# Download the latest version for your architecture. In out case 3.11.4-linux-x86_64
$ curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip

# Unzip the file under $HOME/.local or a directory of your choice. For example:
$ unzip protoc-3.11.4-linux-x86_64.zip -d $HOME/.local

# Update your environment’s path variable to include the path to the protoc executable. For example:
$ export PATH="$PATH:$HOME/.local/bin"
```

### Install the protoc plugin for Go (`protoc-gen-go`)

```bash
# Install the plugin
$ go get github.com/golang/protobuf/protoc-gen-go

# Update your PATH so that the protoc compiler can find the plugin:
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

> In this project we are adding the protoc-gen-plugin to the PATH in the Makefile
