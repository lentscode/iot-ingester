package processor

import (
	"errors"

	"github.com/lentscode/iot-ingester/internal/common"
)

func (p *Processor) insertDataIntoDb(data *common.RawData) error {
	dbData := ProcessorDeviceData{
		RawData: *data,
	}

	result := p.db.Create(&dbData)
	if result.Error != nil {
		err := errors.Join(common.ErrDBInsert, result.Error)
		return err
	}

	return nil
}
