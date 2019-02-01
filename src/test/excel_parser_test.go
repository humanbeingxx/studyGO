package test

import (
	"fmt"
	"testing"
	"tools/excel"
)

func TestExcelParser(t *testing.T) {
	excel := excel.Parse("/Users/xiaoshuang.cui/temp/test.xlsx")

	fmt.Println(excel.Header)
	fmt.Println(excel.Content)
}
