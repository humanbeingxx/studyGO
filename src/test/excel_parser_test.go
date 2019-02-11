package test

import (
	"fmt"
	"studyGO/src/tools/excel"
	"testing"
)

func TestExcelParser(t *testing.T) {
	excel := excel.Parse("/tmp/cxs/test.xlsx")

	fmt.Println(excel.Header)
	fmt.Println(excel.Content)
}
