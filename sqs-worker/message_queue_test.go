package sqsworker

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockedSqsClient struct {
	Response sqs.ReceiveMessageOutput
	sqsiface.SQSAPI
	mock.Mock
}

func (c *mockedSqsClient) ReceiveMessage(input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	c.Called(input)
	return &c.Response, nil
}

func (c *mockedSqsClient) DeleteMessage(input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
	c.Called(input)
	c.Response = sqs.ReceiveMessageOutput{}
	return &sqs.DeleteMessageOutput{}, nil
}

const maxNumberOfMessages = 9
const waitTimeSecond = 11
const visibilityTimeoutSecond = 19
const queueURL = "https://sqs.eu-west-1.amazonaws.com/1234567890/sqs-queue.fifo"

func TestStart(t *testing.T) {

	config := &Config{
		MaxNumberOfMessage: maxNumberOfMessages,
		WaitTimeSecond:     waitTimeSecond,
		VisibilityTimeout:  visibilityTimeoutSecond,
		QueueURL:           queueURL,
	}

	sqsMessage := &sqs.Message{Body: aws.String(`{ "test": "message", "data": "hello" }`)}
	sqsResponse := sqs.ReceiveMessageOutput{Messages: []*sqs.Message{sqsMessage}}
	mockSQSClient := &mockedSqsClient{Response: sqsResponse}

	queue := newJobMessageQueue(config, mockSQSClient)

	ctx, cancel := contextAndCancel()

	t.Run("the queue has correct configuration", func(t *testing.T) {
		assert.Equal(t, queue.config.QueueURL, queueURL, "QueueURL has been set properly")
		assert.Equal(t, queue.config.MaxNumberOfMessage, int64(maxNumberOfMessages), "MaxNumberOfMessage has been set properly")
		assert.Equal(t, queue.config.WaitTimeSecond, int64(waitTimeSecond), "WaitTimeSecond has been set properly")
		assert.Equal(t, queue.config.VisibilityTimeout, int64(visibilityTimeoutSecond), "WaitTimeSecond has been set properly")
	})

	t.Run("the queue has correct default configuration", func(t *testing.T) {
		minimumConfig := &Config{QueueURL: "sqs-queue"}
		worker := newJobMessageQueue(minimumConfig, mockSQSClient)

		assert.Equal(t, worker.config.QueueURL, "sqs-queue", "QueueURL has been set properly")
		assert.Equal(t, worker.config.MaxNumberOfMessage, int64(10), "MaxNumberOfMessage has been set by default")
		assert.Equal(t, worker.config.WaitTimeSecond, int64(20), "WaitTimeSecond has been set by default")
		assert.Equal(t, worker.config.VisibilityTimeout, int64(20), "VisibilityTimeout has been set by default")
	})

	t.Run("Should put a job to the job queue", func(t1 *testing.T) {
		go queue.listen(ctx)
		defer cancel()

		clientParams := buildClientParams()
		deleteInput := &sqs.DeleteMessageInput{QueueUrl: clientParams.QueueUrl}

		t1.Run("the queue successfully sends a message", func(t *testing.T) {
			mockSQSClient.On("ReceiveMessage", clientParams).Return()

			jobQ, ok := <-queue.GetJobs()
			assert.Equal(t, ok, true)
			assert.IsType(t, jobQ, &job{})
			assert.IsType(t, jobQ.Message, sqsMessage)

			mockSQSClient.AssertExpectations(t)
		})

		t1.Run("the queue successfully delete a message", func(t *testing.T) {
			mockSQSClient.On("ReceiveMessage", clientParams).Return()
			mockSQSClient.On("DeleteMessage", deleteInput).Return()

			jobQ, ok := <-queue.GetJobs()
			assert.Equal(t, ok, true)
			err := queue.DeleteMessage(jobQ.Message.ReceiptHandle)

			assert.NoError(t, err)
			mockSQSClient.AssertExpectations(t)
		})
	})

}

func contextAndCancel() (context.Context, context.CancelFunc) {
	delay := time.Now().Add(1 * time.Millisecond)

	return context.WithDeadline(context.Background(), delay)
}

func buildClientParams() *sqs.ReceiveMessageInput {
	url := aws.String(queueURL)

	return &sqs.ReceiveMessageInput{
		QueueUrl:            url,
		MaxNumberOfMessages: aws.Int64(maxNumberOfMessages),
		WaitTimeSeconds:     aws.Int64(waitTimeSecond),
		VisibilityTimeout:   aws.Int64(visibilityTimeoutSecond),
	}
}
