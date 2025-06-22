package processor

import (
	"github.com/segmentio/kafka-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Processor struct {
	rawDataCreatedConsumer *kafka.Reader
	db                     *gorm.DB
}

type ProcessorParams struct {
	KafkaBrokers   []string
	KafkaTopic     string
	KafkaGroupID   string

	DBUrl string
}

func NewProcessor(params *ProcessorParams) (*Processor, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  params.KafkaBrokers,
		MaxBytes: 10e6,
		Topic:    params.KafkaTopic,
		GroupID:  params.KafkaGroupID,
	})

	db, err := gorm.Open(mysql.Open(params.DBUrl))
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&ProcessorDeviceData{})

	return &Processor{
		rawDataCreatedConsumer: reader,
		db:                     db,
	}, nil
}

func (p *Processor) Start() {
	go p.listenForRawData()

	select {}
}
