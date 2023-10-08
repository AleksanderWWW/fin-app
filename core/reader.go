package core

import "fmt"

type Record struct {
	Date  string  `json:"date"`
	Open  float32 `json:"open"`
	Close float32 `json:"close"`
}

type Reader interface {
	FetchData() (map[string][]Record, error)
}

func GetReaderFromProviderString(provider string, initArgs any) (Reader, error) {
	initArgsMap, ok := initArgs.(map[string]interface{})

	if !ok {
		return nil, fmt.Errorf("Malformed arguments %v", initArgs)
	}

	switch provider {
	case "mock":
		return &MockReader{initArgsMap}, nil

	case "stooq":
		return NewStooqReader(initArgsMap)

	default:
		return nil, fmt.Errorf("Invalid data provider - '%s'", provider)
	}
}

var DefaultHeaders = map[string]string{
	"Connection":                "keep-alive",
	"Expires":                   "-1",
	"Upgrade-Insecure-Requests": "1",
	"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
}
