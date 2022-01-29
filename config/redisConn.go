package config

import (
	"errors"
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

// Pool TODO
var Pool *redis.Pool

// InitRedisPool TODO
func InitRedisPool() {
	redisConfName := "redis_config"
	if !C.Viper.IsSet(redisConfName) {
		panic(fmt.Sprintf("connect redis %s failed config not found", redisConfName))
	}
	redisConf := C.Viper.Sub(redisConfName)

	Pool = &redis.Pool{
		MaxActive: redisConf.GetInt("MaxActive"),
		MaxIdle:   redisConf.GetInt("MaxIdle"),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(
				"tcp",
				redisConf.GetString("Address"),
			)
			if err != nil {
				log.Println("Dial failed :", err)
				return nil, err
			}
			if _, err := c.Do("AUTH", redisConf.GetString("Password")); err != nil {
				log.Println("redos auth failed :", err)
				return nil, err
			}
			return c, nil
		},
	}
	conn := Pool.Get()
	defer conn.Close()
	if _, err := conn.Do("ping"); err != nil {
		panic(errors.New("redis conn err:" + err.Error()))
	}
}
