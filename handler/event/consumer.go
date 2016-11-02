package event

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	log "github.com/cihub/seelog"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

//TODO: take optional queue name from command line when starting the app
const (
	queueName         = "event_stream"
	visibilityTimeout = 10
	waitTimeSeconds   = 10
)

// Consumer defines methods to consume events from a queue
type Consumer interface {
	PollForEvents(ctx context.Context)
}

type sqsEventConsumer struct {
	sqs       sqsiface.SQSAPI
	queueURL  string
	processor Processor
}

func NewConsumer(sqs sqsiface.SQSAPI, processor Processor) (Consumer, error) {
	if sqs == nil {
		return nil, errors.Errorf("The SQS API interface is not initialized")
	}
	if processor == nil {
		return nil, errors.Errorf("The event processor is not initialized")
	}

	//TODO: create queue if doesn't exist
	sqsQueueURL, err := getQueueURL(sqs, queueName)
	if err != nil {
		return nil, err
	}

	return &sqsEventConsumer{
		sqs:       sqs,
		queueURL:  sqsQueueURL,
		processor: processor,
	}, nil
}

func getQueueURL(client sqsiface.SQSAPI, queueName string) (string, error) {
	if client == nil {
		return "", errors.Errorf("The SQS client is not initialized")
	}
	if len(queueName) == 0 {
		return "", errors.Errorf("The queue name cannot be empty")
	}

	input := &sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	}
	output, err := client.GetQueueUrl(input)

	if err != nil {
		return "", errors.Wrapf(err, "Could not get the queue url from queue name: %s", queueName)
	} else if output.QueueUrl == nil {
		return "", errors.Errorf("Queue url is empty: %v", output)
	}

	return *output.QueueUrl, nil
}

func (sqsConsumer sqsEventConsumer) PollForEvents(ctx context.Context) {
	log.Infof("Starting to poll for events from SQS")
	for {
		select {
		case <-ctx.Done():
			return
		default:
			sqsConsumer.pollForMessages()
		}
	}
}

func (sqsConsumer sqsEventConsumer) pollForMessages() {
	receiveMessageInput := &sqs.ReceiveMessageInput{
		QueueUrl:          aws.String(sqsConsumer.queueURL),
		VisibilityTimeout: aws.Int64(visibilityTimeout),
		WaitTimeSeconds:   aws.Int64(waitTimeSeconds),
	}

	output, err := sqsConsumer.sqs.ReceiveMessage(receiveMessageInput)
	if err != nil {
		// wrap to get stack trace
		err = errors.Wrap(err, "Could not poll sqs")
		log.Errorf("%+v", err)
		return
	}

	if output == nil || output.Messages == nil {
		log.Debug("Received a blank message from the queue")
		return
	}

	sqsConsumer.processMessages(output.Messages)
}

func (sqsConsumer sqsEventConsumer) processMessages(messages []*sqs.Message) {
	for _, message := range messages {
		err := sqsConsumer.processEvent(message)
		if err != nil {
			log.Errorf("Could not process message: %v: %+v", message, err)
			continue
		}

		err = sqsConsumer.deleteEvent(message)
		if err != nil {
			log.Errorf("Could not delete message %v: %+v", message, err)
		}
	}
}

func (sqsConsumer sqsEventConsumer) processEvent(message *sqs.Message) error {
	if message == nil {
		return errors.Errorf("The sqs message cannot be nil")
	}
	if message.Body == nil {
		return errors.Errorf("The sqs message body cannot be empty")
	}
	return sqsConsumer.processor.ProcessEvent(*message.Body)
}

func (sqsConsumer sqsEventConsumer) deleteEvent(message *sqs.Message) error {
	if message == nil {
		return errors.Errorf("The sqs message cannot be nil")
	}
	if message.ReceiptHandle == nil {
		return errors.Errorf("The sqs message receipt handle cannot be empty")
	}

	deleteMessageInput := &sqs.DeleteMessageInput{
		ReceiptHandle: message.ReceiptHandle,
		QueueUrl:      aws.String(sqsConsumer.queueURL),
	}

	_, err := sqsConsumer.sqs.DeleteMessage(deleteMessageInput)
	if err != nil {
		return errors.Wrap(err, "Could not delete message")
	}

	return nil
}