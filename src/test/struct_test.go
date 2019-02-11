package test

import (
	"fmt"
	"testing"
)

type human struct {
	name string
	age  int
}

type man struct {
	human    human
	strength int
}

func (man man) fight() {
	fmt.Println("For Honor")
}

func TestInit(t *testing.T) {
	oneman := man{
		human: human{
			name: "aman",
			age:  20,
		},
		strength: 10,
	}

	oneman.fight()
}

func TestIP(t *testing.T) {
	fmt.Println([]byte("123.123"))
}

type DataWrap struct {
	Data string
}

func TestDataWrapPointer(t *testing.T) {
	wrap := DataWrap{"data"}

	fmt.Println("print origin ------")
	fmt.Printf("%p\n", &wrap)
	fmt.Printf("%p\n", &wrap.Data)
	printInnerDataPointerByValue(wrap)
	fmt.Printf("origin : %v", wrap)
	printInnerDataPointerByReference(&wrap)
	fmt.Printf("origin : %v", wrap)
}

func printInnerDataPointerByValue(wrap DataWrap) {
	fmt.Println("print func by value ------")
	fmt.Printf("%p\n", &wrap)
	fmt.Printf("%p\n", &wrap.Data)
	wrap.Data = "ModifiedByValue"
}

func printInnerDataPointerByReference(wrap *DataWrap) {
	fmt.Println("print func by reference ------")
	fmt.Printf("%p\n", &wrap)
	fmt.Printf("%p\n", &wrap.Data)
	wrap.Data = "ModifiedByReference"
}
