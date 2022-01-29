package config

const (
	// KeyLength TODO
	KeyLength = 16
	// KeySet TODO
	KeySet = "-"
)

const (
	// QueueName TODO
	QueueName = "delay_queue_queue"
	// NotifyQueueName TODO
	NotifyQueueName = "delay_queue_notify_queue"
	// ZsetName TODO
	ZsetName = "delay_queue_zset"
	// DetectStop TODO
	DetectStop = 10
	// PopQueueTimeout TODO
	PopQueueTimeout = 2 // seconds
)

const (
	// FailPrefix TODO
	FailPrefix = "failCount_"
)
