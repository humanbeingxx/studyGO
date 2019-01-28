package test

import (
	"fmt"
	"testing"
	"tools/csv"
)

func Test(t *testing.T) {
	csvData := csv.Read("/Users/cxs/temp/test.csv")
	fmt.Println(csvData)
}
