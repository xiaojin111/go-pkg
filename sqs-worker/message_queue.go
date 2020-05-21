package sqsworker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

type job struct {
	StartedAt time.Time
	Duration  time.Duration
	Message   *sqs.Message
}

// A messageQueue provides listening to a SQS queue for job messages, and
// providing those job messages as a job channel to workers so the jobs can be
// processed.
type messageQueue struct {
	config    *Config
	jobCh     chan *job
	sqsClient sqsiface.SQSAPI
	log       Logger
}

// newJobMessageQueue creates a new instance of the messageQueue configuring it
// for the SQS service client it will use. The sqsiface.SQSAPI is used so that
// the code could be unit tested in isolating without also testing the SDK.
func newJobMessageQueue(conf *Config, svc sqsiface.SQSAPI) *messageQueue {
	conf.populateDefaultValues()
	return &messageQueue{
		config:    conf,
		jobCh:     make(chan *job, conf.WorkerSize),
		sqsClient: svc,
		log:       conf.Logger,
	}
}

// listen waits for messages to arrive from the SQS queue, parses the JSON
// message and sends the jobs to the job channel to be processed by the worker pool.
func (m *messageQueue) listen(ctx context.Context) {
	m.log.Info("queue: job Message queue starting")
	defer close(m.jobCh)
	defer m.log.Info("queue: job Message queue quitting.")

	for {
		select {
		case <-ctx.Done():
			m.log.Debug("queue: Stopping polling because a context kill signal was sent")
			return
		default:
			msgs, err := m.receiveMsg()
			if err != nil {
				m.log.Error("queue: Failed to read from message queue", err)
				continue
			}

			// Since SQS ReceiveMessage could return multiple messages at once
			// we should loop over then instead of assuming only a single message
			// message is returned. This is also and easier pattern if we want
			// to bump up the number of messages that will be read from SQS at once
			// by default only one message is read.
			for _, msg := range msgs {
				m.jobCh <- &job{
					StartedAt: time.Now(),
					Message:   msg,
				}
			}
		}
	}
}

// receiveMsg reads a message from the SQS job queue. A visibility timeout is set
// so that no other reader will be able to see the message which this service
// received. Preventing duplication of work. And a wait time provides long pooling
// so the service does not need to micro manage its pooling of SQS.
func (m *messageQueue) receiveMsg() ([]*sqs.Message, error) {
	result, err := m.sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(m.config.QueueURL),
		WaitTimeSeconds:     aws.Int64(m.config.WaitTimeSecond),
		VisibilityTimeout:   aws.Int64(m.config.VisibilityTimeout),
		MaxNumberOfMessages: aws.Int64(m.config.MaxNumberOfMessage),
	})
	if err != nil {
		return nil, err
	}
	return result.Messages, nil
}

// DeleteMessage deletes a previously received message from the job message queue
// Once a job is complete it can safely be deleted from the queue so that no
// other service or worker will rerun the job.
func (m *messageQueue) DeleteMessage(receiptHandle *string) error {
	_, err := m.sqsClient.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(m.config.QueueURL),
		ReceiptHandle: receiptHandle,
	})
	return err
}

// GetJobs returns a read only channel to read jobs from. This channel will
// be closed when the messageQueue no longer is listening for further SQS
// job messages.
func (m *messageQueue) GetJobs() <-chan *job {
	return m.jobCh
}
