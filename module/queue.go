package module

import (
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// QueueInit : interface for interacting with queue
type QueueInit interface {
	InitQueue() (*Queue, error)
}

// QueueConfig : Queue configuration
type QueueConfig struct {
	AccessKey       string
	SecretAccessKey string
	SqsURL          string
	AwsRegion       string
}

// InitializeQueue : Use interface to create queueConfig object
func InitializeQueue(init QueueInit) (*Queue, error) {
	queue, err := init.InitQueue()
	if err != nil {
		return nil, err
	}
	return queue, nil
}

// InitQueue : start sqs
func (cfg *QueueConfig) InitQueue() (*Queue, error) {
	credentials := credentials.NewStaticCredentials(cfg.AccessKey, cfg.SecretAccessKey, "")
	awsSession := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(cfg.AwsRegion),
		Credentials: credentials,
	}))

	client := sqs.New(awsSession)
	if client == nil {
		err := errors.New("SQS client not created")
		return nil, err
	}
	queue := &Queue{client, cfg.SqsURL}
	return queue, nil
}

// Queue : all things queue
type Queue struct {
	client   *sqs.SQS
	queueURL string
}

// PollQueue : start SQS polling
func (q *Queue) PollQueue() {
	for {
		params := &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(q.queueURL), // Required
			MaxNumberOfMessages: aws.Int64(1),
			MessageAttributeNames: []*string{
				aws.String("All"), // Required
			},
			WaitTimeSeconds: aws.Int64(10),
		}

		resp, err := q.client.ReceiveMessage(params)
		if err != nil {
			log.Println(err)
			continue
		}
		if len(resp.Messages) > 0 {
			// start goroutine

			// Delete the processed (or invalid) message
			params := &sqs.DeleteMessageInput{
				QueueUrl:      aws.String(q.queueURL),         // Required
				ReceiptHandle: resp.Messages[0].ReceiptHandle, // Required
			}
			_, err = q.client.DeleteMessage(params)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

// GetQueueAttributes : Get all SQS attributes given the queue url
func (q *Queue) GetQueueAttributes() error {

	_, err := q.client.GetQueueAttributes(&sqs.GetQueueAttributesInput{
		QueueUrl: aws.String(q.queueURL)})

	if err != nil {
		return err
	}
	return nil
}
