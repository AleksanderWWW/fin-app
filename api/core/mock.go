package core

import (
	"time"
)

type MockReader struct {
	initArgs map[string]any
}

func (m *MockReader) FetchData() (map[string][]Record, error) {
	return map[string][]Record {
		"symbol": {
			{
				Date:  time.Now().Format("2006-01-02"),
				Open:  21,
				Close: 32,
			},
			{
				Date:  time.Now().AddDate(0, 0, -1).Format("2006-01-02"),
				Open:  31,
				Close: 42,
			},
			{
				Date:  time.Now().AddDate(0, 0, -2).Format("2006-01-02"),
				Open:  34,
				Close: 23,
			},
		},
	}, nil
}
