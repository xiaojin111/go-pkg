package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	sqsworker "gitee.com/jt-heath/go-pkg/v2/sqs-worker"

	"gitee.com/jt-heath/go-pkg/v2/log"
)

const (
	Endpoint  = "http://localhost:4566"
	QueueName = "task-queue"
)

func main() {
	log.SetLevel(log.DebugLevel)

	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("foo", "var", ""),
		Region:      aws.String(endpoints.UsEast1RegionID),
		Endpoint:    aws.String(Endpoint),
	}))
	sqsClient := sqs.New(sess)

	conf := &sqsworker.Config{
		QueueURL:   fmt.Sprintf("%s/queue/%s", Endpoint, QueueName),
		WorkerSize: 3,
	}

	ctx, cancel := context.WithCancel(context.Background())

	worker := sqsworker.NewWorkerPool(ctx, conf, sqsClient)

	worker.Run(sqsworker.HandlerFunc(func(msg *sqs.Message) error {
		// Handler msg
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
