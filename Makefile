# Check if all executables needed are in the PATH before execite any rule/task.
EXECUTABLES = go protoc 
CHECK_EXECUTABLES := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error No $(exec) in PATH. Check README.md for build requirements)))

# Add protoc plugin for Go to the path
export PATH := $(PATH):$(shell go env GOPATH)/bin

all: client server

clean:
	@ rm -f client server helloworld/helloworld.pb.go

client: helloworld/helloworld.pb.go
	@ go build -o client greeter_client/main.go

server: helloworld/helloworld.pb.go
	@ go build -o server greeter_server/main.go 

helloworld/helloworld.pb.go:
	@protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
