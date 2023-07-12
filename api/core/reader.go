package core

import "time"
import "fmt"

type Record struct {
	Date  time.Time `json:"date"`
	Open  float32   `json:"open"`
	Close float32   `json:"close"`
}

type Reader interface {
	FetchData() []Record
}

func GetReaderFromProviderString(provider string, initArgs any) (Reader, error) {
	initArgsMap, ok := initArgs.(map[string]interface{})

	if !ok {
		return nil, fmt.Errorf("Malformed arguments %v", initArgs)
	}

	switch provider {
	case "mock":
		return &MockReader{initArgsMap}, nil
	default:
		return nil, fmt.Errorf("Invalid data provider - '%s'", provider)
	}
}
