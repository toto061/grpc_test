package sensitiveServer

import (
	"context"
	"fmt"
	"grpc/helloworld/proto"
)

type SensitiveServer struct {
	proto.UnimplementedSensitiveFilterServer
}

func (SensitiveServer) Validate(ctx context.Context, in *proto.ValidateRequest) (*proto.ValidateResponse, error) {
	/*out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, SensitiveFilter_Validate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil*/
	fmt.Printf("recv :%v", in)
	return &proto.ValidateResponse{
		Ok:   true,
		Word: "guess",
	}, nil
}
