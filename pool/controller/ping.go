package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	keyWordClient "grpc_test/helloworld/keyword/client"
	"grpc_test/helloworld/proto"
	sensitiveClient "grpc_test/helloworld/sensitive/client"
	"net/http"
)

func Ping(ctx *gin.Context) {
	sPool := sensitiveClient.GetSensitiveClientPool()
	sConn := sPool.Get()
	defer sPool.Put(sConn)
	sensitiveFilterClient := proto.NewSensitiveFilterClient(sConn)
	sIn := &proto.ValidateRequest{Input: "have a nice day"}
	sensitiveValidateResponse, err := sensitiveFilterClient.Validate(context.Background(), sIn)
	fmt.Printf("%+v %+v\n", sensitiveValidateResponse, err)

	kwPool := keyWordClient.GetKeyWordClientPool()
	kwConn := kwPool.Get()
	defer kwPool.Put(kwConn)
	keyWordFilterClient := proto.NewKeyWordFilterClient(sConn)
	kwIn := &proto.MatchRequest{Input: "have a match word"}
	matchResponse, err := keyWordFilterClient.Match(context.Background(), kwIn)
	fmt.Printf("%+v %+v\n", matchResponse, err)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
