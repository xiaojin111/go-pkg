package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jinmukeji/go-pkg/v2/sqs-worker"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("foo", "var", ""),
		Region:      aws.String(endpoints.UsEast1RegionID),
		Endpoint:    aws.String("http://localhost:4576"),
	}))
	sqsClient := sqs.New(sess)

	conf := &sqsworker.Config{
		QueueURL:   "http://localhost:4576/queue/task-queue.fifo",
		WorkerSize: 3,
	}

	ctx, cancel := context.WithCancel(context.Background())

	worker := sqsworker.NewWorkerPool(ctx, conf, sqsClient)

	worker.Run(sqsworker.HandlerFunc(func(msg *sqs.Message) error {
		fmt.Println(aws.StringValue(msg.Body))
		return nil
	}))

	go stopGracefully(cancel)

	worker.WaitForWorkersDone()

}

func stopGracefully(cancel context.CancelFunc) {
	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	fmt.Println("Stopping sqs workers gracefully...")
	cancel()
}
