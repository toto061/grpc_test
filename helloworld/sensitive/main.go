package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpc/helloworld/proto"
	sensitiveServer "grpc/helloworld/sensitive/server"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50055, "")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
		return
	}
	s := grpc.NewServer()
	proto.RegisterSensitiveFilterServer(s, &sensitiveServer.SensitiveServer{})
	log.Printf("server listening as %s\n", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
		return
	}
}
