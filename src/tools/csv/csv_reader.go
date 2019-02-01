package csv

import (
	"encoding/csv"
	"os"
)

type CsvData struct {
	Header  []string
	Content [][]string
}

func Read(path string) *CsvData {
	if len(path) == 0 {
		return nil
	}
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// reader header
	all, err := reader.ReadAll()

	if err != nil {
		return nil
	}

	header := all[0]
	content := all[1:]

	csvData := new(CsvData)
	csvData.Header = header
	csvData.Content = content

	return csvData
}
