package main

import (
	"Net/gRPC/lws"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

/*gRPC server*/

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *lws.String) (*lws.String, error) {
	reply := &lws.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer()
	lws.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
