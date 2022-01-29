package daemon

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/delayQueue/config"
	"github.com/delayQueue/redisOperate"
)

// Detect TODO
func Detect(ctx context.Context, wg *sync.WaitGroup) {
	defer func() {
		if wg != nil {
			log.Println("Detect stopped")
			wg.Done()
		}
	}()
	log.Println("Detector running ...")
	ticker := time.NewTicker(time.Millisecond * 250)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			detect()
		case <-ctx.Done():
			return
		}
	}
}

func detect() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("detect panic,err", r)
		}
	}()

	keysAndScores, err := redisOperate.RangeZSetScore(0, time.Now().Unix())
	if err != nil {
		log.Println("Detector zrangebyscore err:", err)
	}
	if len(keysAndScores) == 0 {
		return
	}

	keysNeedDel := make([]string, 0, len(keysAndScores))
	notifyItems := make([]string, 0)
	items := make([]string, 0)

	for i := 0; i < len(keysAndScores); i += 2 {
		key := keysAndScores[i]
		keysNeedDel = append(keysNeedDel, key)
		if len(keysAndScores[i]) == config.KeyLength {
			items = append(items, key)
		} else {
			notifyItems = append(notifyItems, key)
		}
	}
	err = redisOperate.BatchPushReadyQueue(config.NotifyQueueName, notifyItems)
	if err != nil {
		log.Println("Detector BatchPushReadyQueue err:", err)
	}
	err = redisOperate.BatchPushReadyQueue(config.QueueName, items)
	if err != nil {
		log.Println("Detector BatchPushReadyQueue err:", err)
	}
	log.Println("success push to ready queue ,num:", len(keysNeedDel), time.Now().Unix())
	if err = redisOperate.RemZSet(keysNeedDel); err != nil {
		log.Println("Detector RemZSet err:", err)
	}

}
