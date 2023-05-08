package RPC

// import (
// 	"backend/RPC/pb/servicePb"

// 	"google.golang.org/grpc"
// )

// command protoc --go_out=:. *.proto

// type server struct {
// 	servicePb.UnimplementedPokerServer
//  }

// func SetUp () {
// 	grpcServer := grpc.NewServer()
// 	servicePb.RegisterGreeterServer(grpcServer, servicePb.LoginRequest)
// 	v := servicePb.LoginRequest
// }