package test

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type outer struct {
	inValue []inner
}

type inner struct {
	valueA string
	valueB string
}

func TestValueOf(t *testing.T) {
	outValue := outer{}
	outValue.inValue = make([]inner, 2)
	fmt.Println(outValue)
	reflect.ValueOf(&outValue)
	fmt.Println(outValue)
}

func TestUnitptr(t *testing.T) {
	data := []int{1, 2, 3}
	fmt.Printf("%p\n%p\n%p\n%p\n", &data, &data[0], &data[1], &data[2])
	iptr := uintptr(unsafe.Pointer(&data[0]))
	fmt.Printf("%x\n", iptr)
	addPtr := (*int)(unsafe.Pointer(iptr + 16))
	value := *addPtr
	fmt.Println(value)
}

type myInt int

func TestPointerConvert(t *testing.T) {
	ints := []int{1, 2, 3}

	p := (*[]myInt)(unsafe.Pointer(&ints))

	myInts := *p

	fmt.Println(myInts[1])
}
