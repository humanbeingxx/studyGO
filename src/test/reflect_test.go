package test

import (
	"fmt"
	"reflect"
	"testing"
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
