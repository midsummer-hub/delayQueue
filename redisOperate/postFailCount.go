package redisOperate

import (
	"log"
	"strconv"

	"github.com/delayQueue/config"
	"github.com/garyburd/redigo/redis"
)

// SetFailCount TODO
func SetFailCount(key string, count, timeout int) error {
	conn := config.Pool.Get()
	defer conn.Close()
	countStr := strconv.Itoa(count)
	key = config.FailPrefix + key
	_, err := conn.Do("setex", key, timeout, countStr)
	if err != nil {
		log.Println("redis SetFailCount err :", err)
		return err
	}
	return nil
}

// GetFailCount TODO
func GetFailCount(key string) (int, error) {
	conn := config.Pool.Get()
	defer conn.Close()
	key = config.FailPrefix + key
	count, err := redis.Int(conn.Do("GET", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
			count = 0
			return count, err
		}
		log.Println("redis GetFailCount err:", err)
		return count, err
	}
	return count, nil
}
