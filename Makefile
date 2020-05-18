SERVER_OUT := "bin/server"
CLIENT_OUT := "bin/client"
API_OUT := "pkg/proto/api.pb.go"
PKG := "github.com/atyagi9006/hello-grpc-gw"
SERVER_PKG_BUILD := "${PKG}/server"
CLIENT_PKG_BUILD := "${PKG}/client"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

.PHONY: all api build_server build_client

all: build_server build_client

proto: pkg/proto/api.proto # pkg/proto/api.proto it is just a extra check if proto file is present or not
	@protoc -I pkg/proto/ \
		-I${GOPATH}/src \
		--go_out=plugins=grpc:pkg/proto \
		pkg/proto/api.proto

api: proto ## Auto-generate grpc go sources

dep: ## Get the dependencies
	@go get -v -d ./...

build-server: dep api ## Build the binary file for server
	@go build -i -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

build-client: dep api ## Build the binary file for client
	@go build -i -v -o $(CLIENT_OUT) $(CLIENT_PKG_BUILD)

clean: ## Remove previous builds
	@rm -f $(SERVER_OUT) $(CLIENT_OUT) $(API_OUT)

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
