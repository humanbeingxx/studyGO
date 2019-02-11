package test

import (
	"fmt"
	"testing"
	"tools/excel"
)

func TestExcelParser(t *testing.T) {
	excel := excel.Parse("/tmp/test.xlsx")

	fmt.Println(excel.Header)
	fmt.Println(excel.Content)
}
