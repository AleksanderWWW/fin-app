package core

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AleksanderWWW/fin-app/utils"
)

const stooqBaseUrl string = "https://stooq.com/q/d/l"

var frequenciesAvailable = map[string]bool{
	"d": true,
	"w": true,
	"m": true,
	"q": true,
	"y": true,
}

type StooqDataReader struct {
	symbols   []interface{}
	startDate string
	endDate   string
	freq      string
}

func (s StooqDataReader) getParams(symbol string) map[string]string {
	return map[string]string{
		"s":  symbol,
		"i":  s.freq,
		"d1": strings.Replace(s.startDate, "-", "", -1),
		"d2": strings.Replace(s.endDate, "-", "", -1),
	}
}

func (s StooqDataReader) readSingle(symbol string) ([]Record, error) {
	params := s.getParams(symbol)
	data, err := utils.GetResponse(params, DefaultHeaders, stooqBaseUrl)
	
	if err != nil{
		return []Record{}, err
	}

	dataRows := strings.Split(data, "\n")[1:]
	var dataLen int = len(dataRows)
	records := make([]Record, dataLen)

	for _, row := range dataRows {
		dataSlice := strings.Split(row, ",")
		if len(dataSlice) < 5 {
			break
		}
		open, err := strconv.ParseFloat(dataSlice[1], 32)

		if err != nil {
			continue
		}

		close, err := strconv.ParseFloat(dataSlice[4], 32)
		if err != nil {
			continue
		}
		record := Record{
			Date: dataSlice[0],
			Open: float32(open),
			Close: float32(close),
		}
		records = append(records, record)
	}

	return records, nil
}

func (s StooqDataReader) FetchData() (map[string][]Record, error) {
	var fetchedDataMap = map[string][]Record{}

	if _, ok := frequenciesAvailable[s.freq]; !ok {
		return fetchedDataMap, fmt.Errorf("Frequency '%s' not recognized", s.freq)
	}

	var err error
	for _, symbol := range s.symbols {
		records, err := s.readSingle(s.symbols[0].(string))
		if err != nil {
			continue
		}
		fetchedDataMap[symbol.(string)] = records
	}

	return fetchedDataMap, err
}


func NewStooqReader(initArgsMap map[string]interface{}) (*StooqDataReader, error) {
	symbols, ok := initArgsMap["symbols"].([]interface{})
		if !ok {
			return nil, fmt.Errorf("Missing or malformed param: 'symbols'")
		}

		startDate, ok := initArgsMap["startDate"].(string)
		if !ok {
			return nil, fmt.Errorf("Missing or malformed param: 'startDate'")
		}

		endDate, ok := initArgsMap["endDate"].(string)
		if !ok {
			return nil, fmt.Errorf("Missing or malformed param: 'endDate'")
		}
		freq, ok := initArgsMap["freq"].(string)

		if !ok {
			freq = "d"
		}

		return &StooqDataReader{
			symbols: symbols,
			startDate: startDate,
			endDate: endDate,
			freq: freq,
		}, nil
}
