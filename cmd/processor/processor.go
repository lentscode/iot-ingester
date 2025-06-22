package main

import (
	"log"
	"os"

	"github.com/lentscode/iot-ingester/internal/services/processor"
)

func main() {
	kafkaAddress := os.Getenv("KAFKA_ADDRESS")
	if kafkaAddress == "" {
		panic("kafka address not specified")
	}
	dbUrl := os.Getenv("PROCESSOR_DB_URL")
	if dbUrl == "" {
		panic("processor db url not specified")
	}

	params := processor.ProcessorParams{
		KafkaBrokers: []string{kafkaAddress},
		KafkaTopic:   "raw-data",
		KafkaGroupID: "raw-data-processor",
		DBUrl:        dbUrl,
	}

	processor, err := processor.NewProcessor(&params)
	if err != nil {
		panic(err.Error())
	}

	log.Println("Starting processor")
	processor.Start()
}
