package message

const (
	// StatusUnknown 默认消费状态
	StatusUnknown Status = 0
	// StatusSuccess 消费成功，会从队列中删除消息
	StatusSuccess Status = 1
	// StatusLater 延迟消费，等待下一次消费
	StatusLater Status = 2
)

// Status 消费状态
type Status int
