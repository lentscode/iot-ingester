package main

import (
	"fmt"
	"os"

	"github.com/lentscode/iot-ingester/internal/services/ingester"
)

func main() {
	kafkaAddress := os.Getenv("KAFKA_ADDRESS")
	if kafkaAddress == "" {
		panic("kafka address not specified")
	}

	params := &ingester.IngesterParams{
		Address:       ":8001",
		ProducerTopic: "raw-data",
		KafkaAddress:  os.Getenv("KAFKA_ADDRESS"),
	}

	ingester, err := ingester.NewIngester(params)
	if err != nil {
		panic("can't start ingester")
	}
	fmt.Println("Starting server")

	ingester.Start()
}
