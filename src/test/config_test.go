package test

import (
	"fmt"
	"service"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReadConfig(t *testing.T) {
	configs := service.ReadConfig()
	fmt.Print(configs)
}
func TestAddConfig(t *testing.T) {
	success := service.AddConfig("test_confi1", "test_confi2")

	// Only pass t into top-level Convey calls
	Convey("test add config should succeed", t, func() {
		So(success, ShouldEqual, true)
	})
}

func TestGetConfig(t *testing.T) {
	noValue := service.GetConfig("should_never_exist")
	hasValue := service.GetConfig("name1")
	// Only pass t into top-level Convey calls
	Convey("test get config", t, func() {
		So(noValue, ShouldEqual, "")
		So(hasValue, ShouldEqual, "value1")
	})
}
