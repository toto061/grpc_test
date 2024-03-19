package main

import (
	"github.com/gin-gonic/gin"
	"grpc_test/pool/controller"
)

func main() {
	//sensitiveAddr := "localhost:50055"
	//ketWordAddr := "localhost:50060"
	//sensitivePool, err := services.GetPool(sensitiveAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//keyWordPool, err := services.GetPool(sensitiveAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.Run()
}
