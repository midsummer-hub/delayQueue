package service

import (
	"time"

	"github.com/delayQueue/config"
	"github.com/delayQueue/redisOperate"
	"github.com/delayQueue/utils"
)

// Push TODO
func Push(val, targetAddress string, ttlTime int64) (string, error) {
	id, err := utils.SnowflakeID()
	if err != nil {
		return "", err
	}
	key := id + config.KeySet + targetAddress
	err = redisOperate.SetPayload(key, val)
	if err != nil {
		return "", err
	}
	ttr := time.Now().Unix() + ttlTime
	err = redisOperate.AddZSet(key, ttr)
	if err != nil {
		return "", err
	}
	return key, err
}
