package test

import (
	"fmt"
	"studyGO/src/tools/csv"
	"testing"
)

func Test(t *testing.T) {
	csvData := csv.Read("/Users/cxs/temp/test.csv")
	fmt.Println(csvData)
}
