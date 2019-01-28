package excel

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

type ParsedExcel struct {
	Header  []string
	Content [][]string
}

func Parse(path string) *ParsedExcel {

	file, err := xlsx.OpenFile(path)
	if err != nil {
		fmt.Println("open file error")
		return nil
	}

	oneSheet := file.Sheets[0]
	if len(oneSheet.Rows) == 0 {
		return nil
	}

	colNum := len(oneSheet.Rows[0].Cells)
	header := make([]string, colNum)
	content := make([][]string, len(oneSheet.Rows))

	for index := range content {
		content[index] = make([]string, colNum)
	}

	for line, row := range oneSheet.Rows {
		if line == 0 {
			for index, cell := range row.Cells {
				header[index] = cell.String()
			}
		} else {
			oneLine := content[line]
			for index, cell := range row.Cells {
				oneLine[index] = cell.String()
			}
		}
	}

	parsedExcel := new(ParsedExcel)

	parsedExcel.Header = header
	parsedExcel.Content = content

	return parsedExcel
}
