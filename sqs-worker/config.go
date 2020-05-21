package sqsworker

import (
	"github.com/jinmukeji/go-pkg/v2/log"
)

const defaultMaxNumberOfMessage = 10
const defaultWaitTimeSecond = 20
const defaultVisibilityTimeout = 20

// Config 配置
type Config struct {
	// 每次获取最大的消息数量。默认 10 个
	MaxNumberOfMessage int64

	// SQS Queue URL 地址
	QueueURL string

	// 等待时间，单位秒。默认 20s
	// The duration (in seconds) for which the call waits for a message to arrive in the queue before returning.
	// If a message is available, the call returns sooner than WaitTimeSeconds.
	// If no messages are available and the wait time expires, the call returns successfully with an empty list of messages.
	// https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-short-and-long-polling.html
	WaitTimeSecond int64

	// 消息可见超时时间，单位秒。默认 20s
	// The duration (in seconds) that the received messages are hidden from subsequent retrieve requests after being retrieved by a ReceiveMessage request.
	// https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html
	VisibilityTimeout int64

	// Worker 大小，默认 1
	WorkerSize int64

	// Logger
	Logger Logger
}

func (config *Config) populateDefaultValues() {
	if config.MaxNumberOfMessage == 0 {
		config.MaxNumberOfMessage = defaultMaxNumberOfMessage
	}

	if config.WaitTimeSecond == 0 {
		config.WaitTimeSecond = defaultWaitTimeSecond
	}

	if config.VisibilityTimeout == 0 {
		config.VisibilityTimeout = defaultVisibilityTimeout
	}

	if config.WorkerSize == 0 {
		config.WorkerSize = 1
	}

	if config.Logger == nil {
		config.Logger = log.StandardLogger()
	}
}
