package test

import (
	"fmt"
	"studyGO/src/tools/csv"
	"testing"
)

func Test(t *testing.T) {
	csvData := csv.Read("/tmp/cxs/test.csv")
	fmt.Println(csvData)
}
