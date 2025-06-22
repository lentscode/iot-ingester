package processor

import (
	"context"
	"encoding/json"
	"log"

	"github.com/lentscode/iot-ingester/internal/common"
)

func (p *Processor) listenForRawData() {
	for {
		message, err := p.rawDataCreatedConsumer.ReadMessage(context.Background())
		if err != nil {
			if common.ShouldContinueListening(err) {
				continue
			} else {
				return
			}
		}

		var data common.RawData
		err = json.Unmarshal(message.Value, &data)
		if err != nil {
			log.Println("can't parse incoming data from kafka")
			continue
		}

		err = p.insertDataIntoDb(&data)
		if err != nil {
			log.Println(err.Error())
			continue
		}
	}
}
