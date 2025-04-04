//go:build tools

package generators

import (
	// proto/grpc
	_ "github.com/golang/protobuf/protoc-gen-go"

	// json gateway
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"

	// docs
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"

	// protoc docs
	_ "github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc"
)
