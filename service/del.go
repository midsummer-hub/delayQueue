package service

import "github.com/delayQueue/redisOperate"

// Del TODO
func Del(dataID string) error {
	err := redisOperate.DelPayload(dataID)
	if err != nil {
		return err
	}
	err = redisOperate.RemZSet([]string{dataID})
	if err != nil {
		return err
	}
	return nil
}
