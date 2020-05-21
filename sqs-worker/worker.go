package sqsworker

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

// HandlerFunc is used to define the Handler that is run on for each message
type HandlerFunc func(msg *sqs.Message) error

// HandleMessage wraps a function for handling sqs messages
func (f HandlerFunc) HandleMessage(msg *sqs.Message) error {
	return f(msg)
}

// Handler interface
type Handler interface {
	HandleMessage(msg *sqs.Message) error
}

// A WorkerPool provides a collection of workers, and access to their lifecycle.
type WorkerPool struct {
	ctx     context.Context
	workers []*Worker
	queue   *messageQueue
	log     Logger
	wg      sync.WaitGroup
}

// NewWorkerPool creates a new instance of the worker pool, and creates all the
// workers in the pool. The workers are spun off in their own goroutines and the
// WorkerPool's wait group is used to know when the workers all completed their
// work and existed.
func NewWorkerPool(ctx context.Context, conf *Config, sqsClient sqsiface.SQSAPI) *WorkerPool {

	queue := newJobMessageQueue(conf, sqsClient)

	wp := &WorkerPool{
		ctx:     ctx,
		workers: make([]*Worker, conf.WorkerSize),
		queue:   queue,
		log:     conf.Logger,
	}
	return wp
}

func (w *WorkerPool) Run(handler Handler) {
	//run queue
	go w.queue.listen(w.ctx)

	for i := 0; i < len(w.workers); i++ {
		w.wg.Add(1)
		w.workers[i] = newWorker(i, w.queue, handler, w.log)

		go func(worker *Worker) {
			worker.run()
			w.wg.Done()
		}(w.workers[i])
	}
}

// WaitForWorkersDone waits for the works to of all completed their work and
// exited.
func (w *WorkerPool) WaitForWorkersDone() {
	w.wg.Wait()
}

// A Worker is a individual processor of jobs from the job channel.
type Worker struct {
	id    int
	queue *messageQueue
	hand  Handler
	log   Logger
}

// newWorker creates an initializes a new worker.
func newWorker(id int, queue *messageQueue, handler Handler, logger Logger) *Worker {
	return &Worker{id: id, queue: queue, log: logger, hand: handler}
}

// run reads from the job channel until it is closed and drained.
func (w *Worker) run() {
	w.log.Info(fmt.Sprintf("worker %d: starting", w.id))
	defer w.log.Info(fmt.Sprintf("worker %d: quitting.", w.id))

	for {
		job, ok := <-w.queue.GetJobs()
		if !ok {
			return
		}
		w.log.Debug(fmt.Sprintf("worker %d: getting message from queue: %s", w.id, aws.StringValue(job.Message.MessageId)))

		err := w.processJob(job, w.hand)
		if err != nil {
			w.log.Error(err)
		}
	}
}

func (w *Worker) processJob(job *job, h Handler) error {
	err := h.HandleMessage(job.Message)
	if err != nil {
		w.log.Error(err)
	}
	err = w.queue.DeleteMessage(job.Message.ReceiptHandle)
	if err != nil {
		return err
	}
	w.log.Debug(fmt.Sprintf("worker %d: deleted message from queue: %s", w.id, aws.StringValue(job.Message.MessageId)))

	job.Duration = time.Since(job.StartedAt)
	w.log.Debug(fmt.Sprintf("worker %d: processed job in: %s", w.id, job.Duration))
	return nil
}
