package common

import "github.com/segmentio/kafka-go"

func NewWriter(topic, address string) *kafka.Writer {
	return &kafka.Writer{
		Topic: topic,
		Addr: kafka.TCP(address),
		Balancer: &kafka.LeastBytes{},
	}
}