package api

import (
	"context"
	"errors"

	"github.com/delayQueue/queueRpc"
	"github.com/delayQueue/service"
)

// Del TODO
func (Server) Del(_ context.Context, req *queueRpc.DelRequest) (*queueRpc.DelResponse, error) {
	if req.DataID == "" {
		return nil, errors.New("data to be deleted must have a key")
	}
	err := service.Del(req.DataID)
	if err != nil {
		return nil, err
	}
	return &queueRpc.DelResponse{}, nil
}
