package core

import "time"

type StooqDataReader struct {
	symbols   []string
	startDate time.Time
	endDate   time.Time
	freq      string
	baseUrl   string
}
