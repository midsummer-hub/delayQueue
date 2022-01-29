package api

import (
	"context"
	"errors"

	"github.com/delayQueue/queueRpc"
	"github.com/delayQueue/service"
)

// Pop TODO
func (Server) Pop(_ context.Context, req *queueRpc.PopRequest) (*queueRpc.PopResponse, error) {

	if req.ExpiryTime < 0 {
		return nil, errors.New("the expiration time is greater than 0")
	}
	data, err := service.Pop(req.ExpiryTime)
	if err != nil {
		return nil, err
	}
	return &queueRpc.PopResponse{Data: data}, nil
}
