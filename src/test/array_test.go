package test

import (
	"fmt"
	"reflect"
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

func TestSliceType(t *testing.T) {
	ar := [3]int{1, 2, 3}
	sl := []int{1, 2, 3}

	fmt.Println(reflect.TypeOf(ar))
	fmt.Println(reflect.TypeOf(sl))
}

func TestAppend(t *testing.T) {
	slice := []int{1}

	a1 := append(slice, 2)
	fmt.Println(slice)
	fmt.Println(a1)

	a2 := append(slice, 3)
	fmt.Println(slice)
	fmt.Println(a2)
	fmt.Println(a1)

	fmt.Println(&slice[0])
	fmt.Println(&a1[0])
	fmt.Println(&a2[0])

	fmt.Printf("%d", reflect.TypeOf(&slice[0]))
}
