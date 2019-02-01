package test

import (
	"fmt"
	"studyGO/src/tools/excel"
	"testing"
)

func TestExcelParser(t *testing.T) {
	excel := excel.Parse("/Users/xiaoshuang.cui/temp/test.xlsx")

	fmt.Println(excel.Header)
	fmt.Println(excel.Content)
}
