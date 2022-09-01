package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-redis/redis"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		const redisKey1 = "test"
		const redisKey2 = "test2"
		var duration time.Duration
		var err error

		client := redis.NewClient(&redis.Options{
			Addr: "redis:6379",
		})

		err = client.Set(
			redisKey1, fmt.Sprintf("id: %v\nbody: %v\n", message.MessageId, message.Body), time.Duration(6)*time.Second,
		).Err()

		if err != nil {
			return err
		}

		time.Sleep(time.Duration(5) * time.Second)

		duration, err = client.TTL(redisKey1).Result()

		if err != nil {
			return err
		}

		err = client.Set(redisKey2, fmt.Sprintf("%v", duration.Milliseconds()), 0).Err()

		if err != nil {
			return err
		}
	}

	return nil
}
