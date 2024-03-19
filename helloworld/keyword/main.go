package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	keywordServer "grpc_test/helloworld/keyword/server"
	"grpc_test/helloworld/proto"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50060, "")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
		return
	}
	s := grpc.NewServer()
	proto.RegisterKeyWordFilterServer(s, &keywordServer.KeywordServer{})
	log.Printf("server listening as %s\n", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
		return
	}
}
