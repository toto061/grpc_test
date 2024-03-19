package client

import (
	services "grpc_test/pool/service"
	"log"
)

type ServiceClient interface {
	GetPool(add string) services.ClientPool
}

type DefaultClient struct {
}

func (c *DefaultClient) GetPool(addr string) services.ClientPool {
	pool, err := services.GetPool(addr, c.getOptions()...)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return pool
}
