package common

type RawData struct {
	DeviceID string  `json:"device_id"`
	Tag      string  `json:"tag"`
	Value    float64 `json:"value"`
}

func (r *RawData) IsValid() bool {
	return r.DeviceID != "" && r.Tag != ""
}
