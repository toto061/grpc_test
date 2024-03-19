package keyWordClient

import (
	services "grpc_test/pool/service"
	"grpc_test/pool/service/client"
	"sync"
)

type keyWordClient struct {
	client.DefaultClient
}

var pool services.ClientPool
var once sync.Once
var keyWordAddr = "localhost:50060"

func GetKeyWordClientPool() services.ClientPool {
	once.Do(func() {
		c := &keyWordClient{}
		pool = c.GetPool(keyWordAddr)

	})
	return pool
}
