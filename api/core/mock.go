package core

import (
	"time"
)

type MockReader struct{
	initArgs map[string]any
}

func (m *MockReader) FetchData() []Record {
	return []Record{
		{
			Date: time.Now(),
			Open: 21,
			Close: 32,
		},
		{
			Date: time.Now().AddDate(0, 0, -1),
			Open: 31,
			Close: 42,
		},
		{
			Date: time.Now().AddDate(0, 0, -2),
			Open: 34,
			Close: 23,
		},
	}
}
