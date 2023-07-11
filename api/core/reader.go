package core

import "time"


type Record struct {
	Date time.Time `json:"date"`
	Open float32 `json:"open"`
	Close float32 `json:"close"`
}

type Reader interface {
	FetchData() []Record
}
