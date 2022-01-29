package service

import (
	"github.com/delayQueue/config"
	"github.com/delayQueue/redisOperate"
)

// Pop TODO
func Pop(ttlTime int64) (string, error) {
	key, err := redisOperate.PopReadyQueue(config.QueueName, int(ttlTime))
	if err != nil {
		return "", err
	}
	data, err := redisOperate.GetPayload(key)
	if err != nil {
		return "", err
	}
	return data, nil
}
