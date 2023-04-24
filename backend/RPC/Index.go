package RPC

import (
	"backend/RPC/RPC_Build"
	"backend/RPC/RPC_Origin"

	"google.golang.org/grpc"

)

// command protoc --go_out=:. *.proto

func SetUp () {
	grpcServer := grpc.NewServer()

	RPC_Origin.RegisterGreeterServer(grpcServer, RPC_Build.LoginRequest)
}