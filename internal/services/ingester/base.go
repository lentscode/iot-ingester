package ingester

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lentscode/iot-ingester/internal/common"
	"github.com/segmentio/kafka-go"
)

type Ingester struct {
	address                string
	router                 *mux.Router
	rawDataCreatedProducer *kafka.Writer
}

type IngesterParams struct {
	Address       string
	ProducerTopic string
	KafkaAddress  string
}

func NewIngester(params *IngesterParams) (*Ingester, error) {
	return &Ingester{
		address:                params.Address,
		router:                 mux.NewRouter(),
		rawDataCreatedProducer: common.NewWriter(params.ProducerTopic, params.KafkaAddress),
	}, nil
}

func (i *Ingester) Start() error {
	i.router.HandleFunc("/raw-data", i.handleRawData)

	err := http.ListenAndServe(i.address, i.router)
	if err != nil {
		return err
	}

	return nil
}
