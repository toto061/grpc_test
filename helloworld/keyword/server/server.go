package keywordServer

import (
	"context"
	"fmt"
	"grpc_test/helloworld/proto"
)

type KeywordServer struct {
	proto.UnimplementedKeyWordFilterServer
}

func (KeywordServer) Match(ctx context.Context, in *proto.MatchRequest) (*proto.MatchResponse, error) {
	fmt.Printf("recv :%v", in)
	return &proto.MatchResponse{
		Ok:   true,
		Word: "match",
	}, nil
}
