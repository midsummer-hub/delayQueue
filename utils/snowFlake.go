package utils

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

const (
	workerBites uint8 = 10                        // 每台机器（节点）ID位数，10位最大可以有2^10=1024个节点
	numberBites uint8 = 12                        // 每个集群下的每个节点，1毫秒内可生成的ID序号的二进制,每秒可生成2^12-1=4049个唯一ID
	workerMax   int64 = -1 ^ (-1 << workerBites)  // 节点ID最大值，防止溢出
	numberMax   int64 = -1 ^ (-1 << numberBites)  // ID序号的最大值，防止溢出
	timeShift         = workerBites + numberBites // 时间戳想左偏移量
	workerShift       = numberBites               // 节点ID向左的偏移量
	epoch       int64 = 1643342400000             // 一旦定义且开始生成ID千万不要修改，不然可能会生成相同的ID
)

// Worker 定义工作节点基本参数
type Worker struct {
	mu        sync.Mutex // 添加互斥锁，保证并发安全
	timestamp int64      // 记录时间戳
	workerID  int64      // 该节点ID
	number    int64      // 当前毫秒已经生成的iD序列号，（从0开始累加），1毫秒内最多生成4096个ID
}

// ID TODO
type ID int64

// NewWorker 实例化工作节点
func NewWorker(workerID int64) (*Worker, error) {
	// 检查workerID是否在在定义的范围内
	if workerID < 0 || workerID > workerMax {
		return nil, errors.New("workerID number must be between 0 and " + string(workerMax))
	}
	return &Worker{
		timestamp: 0,
		workerID:  workerID,
		number:    0,
	}, nil
}

// GetID 获取ID
func (w *Worker) GetID() ID {
	w.mu.Lock()

	now := time.Now().UnixNano() / 1e6 // 纳秒转毫秒
	if w.timestamp == now {
		w.number++
		// 判断当前节点1秒内已经生成的ID是否超出设定的numberMax范围
		if w.number > numberMax {
			// 如果当前节点在1秒内生成的节点超出上限，需等待1毫秒继续生成
			for now < w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		// 如果当前时间与工作节点上一次生成的ID的时间不一致，则需要重置工作节点生成ID的序号
		w.number = 0
		w.timestamp = now
	}

	id := (now-epoch)<<timeShift | (w.workerID << workerShift) | (w.number)
	w.mu.Unlock()
	return ID(id)
}

// String 转换为字符串
func (f ID) String() string {
	return strconv.FormatInt(int64(f), 10)
}
