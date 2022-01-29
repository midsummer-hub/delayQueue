package api

import (
	"context"
	"log"

	"github.com/delayQueue/queueRpc"
)

// Server TODO
type Server struct {
	queueRpc.UnimplementedDelayQueueServer
}

// Ping TODO
func (Server) Ping(_ context.Context, req *queueRpc.PingRequest) (*queueRpc.PingResponse, error) {
	log.Println("recv ping msg :", req.Msg)
	return &queueRpc.PingResponse{Msg: "ping"}, nil
}
