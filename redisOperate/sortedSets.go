package redisOperate

import (
	"github.com/delayQueue/config"
	"github.com/garyburd/redigo/redis"
)

// AddZSet TODO
func AddZSet(key string, score int64) (err error) {
	conn := config.Pool.Get()
	defer config.Pool.Close()
	_, err = conn.Do("ZADD", config.ZsetName, score, key)
	if err != nil {
		return
	}
	return
}

// RemZSet TODO
func RemZSet(keys []string) (err error) {
	if len(keys) == 0 {
		return
	}
	conn := config.Pool.Get()
	defer config.Pool.Close()
	_, err = conn.Do("ZREM", redis.Args{}.Add(config.ZsetName).AddFlat(keys)...)
	return
}

// RangeZSet TODO
func RangeZSet(start, end int) (keys []string, err error) {
	conn := config.Pool.Get()
	defer config.Pool.Close()
	keys, err = redis.Strings(conn.Do("ZRANGE", config.ZsetName, start, end, "withscores"))
	return
}

// RangeZSetScore TODO
func RangeZSetScore(start, end int64) (keys []string, err error) {
	conn := config.Pool.Get()
	defer conn.Close()
	keys, err = redis.Strings(conn.Do("ZRANGEBYSCORE", config.ZsetName, start, end, "withscores"))
	return
}
