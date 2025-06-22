package processor

import (
	"github.com/lentscode/iot-ingester/internal/common"
	"gorm.io/gorm"
)

type ProcessorDeviceData struct {
	gorm.Model

	common.RawData
}
