package sensitiveClient

import (
	services "grpc_test/pool/service"
	"grpc_test/pool/service/client"
	"sync"
)

type sensitiveClient struct {
	client.DefaultClient
}

var pool services.ClientPool
var once sync.Once
var sensitiveAddr = "localhost:50055"

func GetSensitiveClientPool() services.ClientPool {
	once.Do(func() {
		c := &sensitiveClient{}
		pool = c.GetPool(sensitiveAddr)

	})
	return pool
}
