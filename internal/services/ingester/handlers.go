package ingester

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/lentscode/iot-ingester/internal/common"
	"github.com/segmentio/kafka-go"
)

// "/raw-data"
func (i *Ingester) handleRawData(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var data RawData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		err = errors.Join(common.ErrReadingRequest, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !data.isValid() {
		http.Error(w, common.ErrRequestSchema.Error(), http.StatusBadRequest)
		return
	}

	dataToSend, err := json.Marshal(data)
	if err != nil {
		err = errors.Join(errors.New("can't convert into bytes"), err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := kafka.Message{
		Key:   []byte(data.DeviceID),
		Value: dataToSend,
	}

	err = i.rawDataCreatedProducer.WriteMessages(ctx, message)
	if err != nil {
		err = errors.Join(errors.New("can't create event"), err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
