package rabittmq

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Queue struct {
	Name string
}

func NewQueue(name string) *Queue {
	return &Queue{Name: name}
}

func (q *Queue) Publish(msg any) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return fmt.Errorf("Failed to connect to rabbitmq")
	}
	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("Failed to open channel")
	}
	queue, err := ch.QueueDeclare(
		q.Name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Failed to declare queue")
	}

	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("Faield to parse message")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return fmt.Errorf("Failed to publish message")
	}

	return nil
}
