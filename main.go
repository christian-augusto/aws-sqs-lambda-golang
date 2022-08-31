package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-redis/redis"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		var err error

		client := redis.NewClient(&redis.Options{
			Addr: "redis:6379",
		})

		err = client.Set("test", fmt.Sprintf("id: %v\nbody: %v\n", message.MessageId, message.Body), 0).Err()

		if err != nil {
			return err
		}
	}

	return nil
}
