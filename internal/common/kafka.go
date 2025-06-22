package common

import (
	"context"
	"errors"
	"io"
	"log"

	"github.com/segmentio/kafka-go"
)

func NewWriter(topic, address string) *kafka.Writer {
	return &kafka.Writer{
		Topic:    topic,
		Addr:     kafka.TCP(address),
		Balancer: &kafka.LeastBytes{},
	}
}

func ShouldContinueListening(err error) bool {
	if errors.Is(err, io.EOF) {
		log.Println("no message available or reader closed")
		return true
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		log.Println("context canceled or deadline exceeded")
		return true
	}
	return false
}