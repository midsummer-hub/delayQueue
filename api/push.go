package api

import (
	"context"
	"errors"

	"github.com/delayQueue/queueRpc"
	"github.com/delayQueue/service"
)

// Push TODO
func (Server) Push(_ context.Context, req *queueRpc.PushRequest) (*queueRpc.PushResponse, error) {

	if req.Data == "" || req.ExpiryTime < 0 {
		return nil, errors.New("params error")
	}
	dataId, err := service.Push(req.Data, req.TargetAddress, req.ExpiryTime)
	if err != nil {
		return nil, err
	}
	return &queueRpc.PushResponse{DataID: dataId}, nil
}
