package test

import (
	"fmt"
	"testing"
)

func TestMultiArray(t *testing.T) {
	x := 2
	y := 3

	table := make([][]int, x)

	for i := range table {
		table[i] = make([]int, y)
	}

	fmt.Println(table)
}

func TestMultiSlice(t *testing.T) {
	x := 2
	y := 3

	data := make([]int, x*y)

	for i := range data {
		data[i] = i
	}

	table := make([][]int, x)

	for i := range table {
		table[i] = data[i*y : (i+1)*y]
	}

	fmt.Println(table, &data[0], &table[0][0])
}

func TestStringSlice(*testing.T) {
	str := "abc"

	bytes := []byte(str)
	bytes[0] = 'A'

	fmt.Println(string(bytes))
	fmt.Println(str)

	fmt.Println(rune(bytes[0]))
	fmt.Println('a')
}
