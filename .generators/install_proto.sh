#!/bin/bash
# Script for build proto file grpc service and model

DIR_PROTO_BIN="./.generators/bin" 
DIR_GENERATORS="./.generators"
GO_PATH=$HOME"/go"
GO_BIN=$GO_PATH"/bin"
GO_BIN_PROTOC_GEN_GO=$GO_BIN"/protoc-gen-go" 
GO_BIN_PROTOC_GEN_GO_GRPC=$GO_BIN"/protoc-gen-go-grpc"
GO_BIN_PROTOC_GEN_SWAGGER=$GO_BIN"/protoc-gen-openapiv2"
GO_BIN_PROTOC_GEN_GRPC_GATEWAY=$GO_BIN"/protoc-gen-grpc-gateway"

if (( $EUID == 0 )); then
    echo "Please don't run as root"
    exit
fi

if [ ! -d "$DIR_GENERATORS" ]; then
  echo "run script from root of project dir"
  exit
fi

# create bin directory
if [ ! -d "$DIR_PROTO_BIN" ]; then
  mkdir -p $DIR_PROTO_BIN
fi

export GO111MODULE=off
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-openapiv2
go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-openapiv2
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

export GO111MODULE=on

go mod tidy
go mod vendor

cp "$GO_BIN_PROTOC_GEN_GO" "$DIR_PROTO_BIN" 
cp "$GO_BIN_PROTOC_GEN_GO_GRPC" "$DIR_PROTO_BIN" 
cp "$GO_BIN_PROTOC_GEN_SWAGGER" "$DIR_PROTO_BIN" 
cp "$GO_BIN_PROTOC_GEN_GRPC_GATEWAY" "$DIR_PROTO_BIN" 


