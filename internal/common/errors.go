package common

import "errors"

var (
	ErrRequestSchema  = errors.New("request does not conform to schema")
	ErrReadingRequest = errors.New("error reading request")
)
