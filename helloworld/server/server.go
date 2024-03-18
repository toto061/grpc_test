package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc/helloworld/proto"
	"io"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "")
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {
	log.Printf("server recv: %v\n", in)
	return &proto.HelloResponse{
		Msg: "hello client---",
	}, nil
}

func (server) SayHelloClientStream(stream proto.Greeter_SayHelloClientStreamServer) error {
	var i = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.HelloResponse{Msg: fmt.Sprintf("server closed.recv count:%d", i)})
		}
		fmt.Printf("recv : %v\n", req)
		i++

	}
	return status.Errorf(codes.Unimplemented, "method SayHelloClientStream not implemented")
}
func (server) SayHelloServerStream(n *proto.HelloRequest, stream proto.Greeter_SayHelloServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloServerStream not implemented")
}
func (server) SayHelloTwoWayStream(stream proto.Greeter_SayHelloTwoWayStreamServer) error {
	var i = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("server recv err: %v\n", err)
		}
		fmt.Printf("recv : %v\n", req)
		i++
		stream.Send(&proto.HelloResponse{Msg: fmt.Sprintf("respone :%d", i)})
	}
	return nil //status.Errorf(codes.Unimplemented, "method SayHelloTwoWayStream not implemented")
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
		return
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	log.Printf("server listening as %s\n", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
		return
	}
}
