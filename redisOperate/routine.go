package redisOperate

import (
	"github.com/delayQueue/config"
	"github.com/delayQueue/module"
	"github.com/garyburd/redigo/redis"
)

// SetPayload TODO
func SetPayload(key, value string) (err error) {
	conn := config.Pool.Get()
	defer conn.Close()
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return
	}
	return
}

// GetPayload TODO
func GetPayload(key string) (payload string, err error) {
	conn := config.Pool.Get()
	defer conn.Close()
	payload, err = redis.String(conn.Do("GET", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
			return
		}
		return
	}
	return
}

// DelPayload TODO
func DelPayload(key string) (err error) {
	conn := config.Pool.Get()
	defer conn.Close()
	_, err = conn.Do("DEL", key)
	return
}

// BatchSetPayload TODO
func BatchSetPayload(kv []*module.RedisKV) (successCount int, err error) {
	conn := config.Pool.Get()
	defer conn.Close()
	for _, val := range kv {
		err := SetPayload(val.Key, val.Value)
		if err != nil {
			continue
		}
		successCount++
	}
	return
}
