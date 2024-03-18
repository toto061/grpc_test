package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc/helloworld/proto"
	"io"
	"log"
	"time"
)

var (
	addr = flag.String("addr", "localhost:50051", "")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	//sayHello(c)

	//sayHelloClientStream(c)

	sayHelloTwoWayStream(c)
}

func getHelloRequest() *proto.HelloRequest {
	birthday := timestamppb.New(time.Now())
	any1, _ := anypb.New(birthday)
	in := &proto.HelloRequest{
		Name:     "jkljlkjlk",
		Age:      18,
		Gender:   proto.Gender_MALE,
		Birthday: birthday,
		Hobys:    []string{"badminton", "football", "reading"},
		Addr: &proto.Address{
			Province: "hb",
			City:     "wuhan",
		},
		Data: map[string]*anypb.Any{
			"a": any1,
		},
	}
	return in
}

func sayHello(c proto.GreeterClient) {
	ctx := context.Background()

	r, err := c.SayHello(ctx, getHelloRequest())
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Fatal(r.Msg)
}

func sayHelloClientStream(c proto.GreeterClient) {
	ctx := context.Background()
	list := []*proto.HelloRequest{
		getHelloRequest(), getHelloRequest(), getHelloRequest(),
	}

	stream, err := c.SayHelloClientStream(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, in := range list {
		err := stream.Send(in)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("client recv:%v\n", response)
}

func sayHelloTwoWayStream(c proto.GreeterClient) {
	ctx := context.Background()
	list := []*proto.HelloRequest{
		getHelloRequest(), getHelloRequest(), getHelloRequest(),
	}

	stream, err := c.SayHelloTwoWayStream(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	var done = make(chan struct{}, 0)
	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatal(err)
				close(done)
				return
			}
			fmt.Printf("client recv:%v\n", response)
		}
	}()

	for _, in := range list {
		err := stream.Send(in)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	stream.CloseSend()
	<-done

}
