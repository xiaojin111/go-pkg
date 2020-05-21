# sqs-worker

> Forked from: https://github.com/Stashchenko/sqs-worker

Worker service which reads from a SQS queue pulls off job messages and processes them concurrently.

Required config fields:

- **QueueURL** - QueueUrl is a queue url, ex. http://localhost:4576/queue/task-queue.fifo

Optionally the follow config variables can be provided.

- **MaxNumberOfMessage** - SQS ReceiveMessage could return multiple messages at once This is also and easier pattern if we want to bump up the number of messages that will be read from SQS at once by default 10 messages are read.
- **WaitTimeSecond** - The duration (in seconds) for which the call waits for a message to arrive in the queue before returning. If a message is available, the call returns sooner than WaitTimeSeconds.
- **VisibilityTimeout** - The duration (in seconds) that the received messages are hidden from subsequent retrieve requests after being retrieved by a ReceiveMessage request.
- **WorkerSize** - Numbers of max parallel workers
- **Logger** - by default used stdout. Supports [logrus](https://github.com/sirupsen/logrus) interface

## Usage

```go
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
	signal.Notify(signalChan, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-signalChan
	fmt.Println("Stopping sqs workers gracefully...")
	cancel()
}
```

### Run SQS Servce (localstack) from DockerHub:

```sh
docker run --rm \
 	-p 4576:4576 \
 	-e SERVICES:sqs \
 	-d --name sqs \
 	localstack/localstack
```

#### Create the SQS Queue:

```sh
aws --endpoint http://localhost:4576 \
	sqs create-queue \
	--queue-name task-queue.fifo \
	--attributes '{"FifoQueue": "true", "ContentBasedDeduplication":"true"}'
```

#### Check the SQS Queue:

```sh
aws --endpoint http://localhost:4576 sqs list-queues
```

#### Add messages to the SQS Queue:

```sh
aws --endpoint http://localhost:4576 \
	sqs send-message \
	--queue-url http://localhost:4576/queue/task-queue.fifo \
	--message-body '{"msg": "hello"}' \
	--message-group-id "a-string"
```

or use `sh ./gen_messages.sh` to generate 100 messages and send it to a queue

##### Example of gracefully shutdown

```bash
    2020/05/20 15:02:36 [INFO] queue: job Message queue starting
    2020/05/20 15:02:36 [INFO] worker 0: starting
    2020/05/20 15:02:36 [INFO] worker 2: starting
    2020/05/20 15:02:36 [INFO] worker 1: starting
    ...
    2020/05/20 15:03:13 [DEBUG] worker 0: getting message from queue: 7a29e48e-cf9e-4e6b-a9aa-f31e6600f5e0
    {'msg': 'hello46'}
    2020/05/20 15:03:13 [DEBUG] worker 0: deleted message from queue: 7a29e48e-cf9e-4e6b-a9aa-f31e6600f5e0
    2020/05/20 15:03:13 [DEBUG] worker 0: processed job in: 9.999398ms
    2020/05/20 15:03:14 [DEBUG] worker 2: getting message from queue: 16d0df57-4a77-4a88-b83a-8f5eee2c4312
    {'msg': 'hello47'}
    2020/05/20 15:03:14 [DEBUG] worker 2: deleted message from queue: 16d0df57-4a77-4a88-b83a-8f5eee2c4312
    2020/05/20 15:03:14 [DEBUG] worker 2: processed job in: 12.541419ms

    #Stop via ^C

    Stopping sqs workers gracefully...
    2020/05/20 15:03:14 [DEBUG] queue: Stopping polling because a context kill signal was sent
    2020/05/20 15:03:14 [INFO] queue: job Message queue quitting.
    2020/05/20 15:03:14 [INFO] worker 0: quitting.
    2020/05/20 15:03:14 [INFO] worker 2: quitting.
    2020/05/20 15:03:14 [DEBUG] worker 1: getting message from queue: 5b1a52b0-34b6-4b63-9eb5-96d70db355c0
    {'msg': 'hello48'}
    2020/05/20 15:03:15 [DEBUG] worker 1: deleted message from queue: 5b1a52b0-34b6-4b63-9eb5-96d70db355c0
    2020/05/20 15:03:15 [DEBUG] worker 1: processed job in: 11.258625ms
    2020/05/20 15:03:15 [INFO] worker 1: quitting.

    Process finished with exit code 0
```
