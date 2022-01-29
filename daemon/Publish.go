package daemon

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/delayQueue/config"
	"github.com/delayQueue/httpClient"
	"github.com/delayQueue/redisOperate"
)

// Publish TODO
func Publish(ctx context.Context, wg *sync.WaitGroup) {
	defer func() {
		if wg != nil {
			log.Println("publish stopped")
			wg.Done()
		}
	}()
	log.Println("publish running...")
	for {
		select {
		case <-ctx.Done():
			return
		default:
			handler()
		}
	}

}
func handler() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("publish handler panic err:", r)
		}
	}()
	key, err := redisOperate.PopReadyQueue(config.NotifyQueueName, config.PopQueueTimeout)
	if err != nil {
		log.Println("publish pop ready queue err:", err)
	}

	payload, err := redisOperate.GetPayload(key)
	if err != nil {
		log.Println("publisher GetPayload err:", err)
		err = redisOperate.PushReadyQueue(config.NotifyQueueName, key)
		if err != nil {
			log.Println("GetPayload err And PushReadyQueue err:", err)
		}
		return
	}

	if len(payload) == 0 {
		return
	}

	go publish(key, payload)
}

// PostHeartbeat TODO
var PostHeartbeat = []int{0, 2, 8, 30, 60 * 2, 60 * 5, 60 * 30, 60 * 60}

func publish(key, payload string) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("publisher panic err:", r)
		}
	}()
	url := key[config.KeyLength+1:]
	err := httpClient.SendPostRequest(url, payload)
	if err != nil {
		count, err := handlePostErr(key)
		log.Println("Publish send post count:", count, "error:", err, time.Now())
		if err != nil {
			log.Println("rePost fail", err)
		}
		return
	}
	log.Println("post success", key, time.Now())
	_ = redisOperate.DelPayload(key)
}

func handlePostErr(key string) (count int, err error) {

	count, _ = redisOperate.GetFailCount(key)
	count++
	if count >= len(PostHeartbeat) {
		return
	}
	delay := PostHeartbeat[count]
	nextPostTime := time.Now().Unix() + int64(delay)
	err = redisOperate.SetFailCount(key, count, delay<<1)
	if err != nil {

	}
	err = redisOperate.AddZSet(key, nextPostTime)
	if err != nil {
		return
	}
	return
}
