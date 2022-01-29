package redisOperate

import (
	"github.com/delayQueue/config"
	"github.com/garyburd/redigo/redis"
)

// PushReadyQueue TODO
func PushReadyQueue(queueName string, key string) (err error) {
	conn := config.Pool.Get()
	defer conn.Close()
	_, err = conn.Do("LPUSH", queueName, key)
	if err != nil {
		return
	}
	return
}

// BatchPushReadyQueue TODO
func BatchPushReadyQueue(queueName string, keys []string) (err error) {
	conn := config.Pool.Get()
	defer conn.Close()
	_, err = conn.Do("LPUSH", redis.Args{}.Add(queueName).AddFlat(keys)...)
	if err != nil {
		return
	}
	return
}

// PopReadyQueue TODO
func PopReadyQueue(queueName string, timeout int) (key string, err error) {
	conn := config.Pool.Get()
	defer conn.Close()
	nameData, err := redis.Strings(conn.Do("BRPOP", queueName, timeout))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
			return
		}
		return
	}
	if len(nameData) > 1 {
		key = nameData[1]
	}
	return
}
