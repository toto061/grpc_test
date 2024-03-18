package poolService

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"sync"
)

type clientPool struct {
	pool sync.Pool
}

func (c *clientPool) Get() *grpc.ClientConn {
	conn := c.pool.Get().(*grpc.ClientConn)
	if conn.GetState() == connectivity.Shutdown || conn.GetState() == connectivity.TransientFailure {

	}
	return nil
}

func (c *clientPool) Put(conn *grpc.ClientConn) {

}
