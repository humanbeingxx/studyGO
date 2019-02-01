package test

import (
	"fmt"
	"math/rand"
	"service"
	"strconv"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRandInt(t *testing.T) {
	ran := rand.Intn(100)

	fmt.Println(ran)
}

func TestNow(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Unix())
}

func TestReadConfig(t *testing.T) {
	configs := service.ReadConfig()
	fmt.Print(configs)
}

func TestAddConfig(t *testing.T) {
	name := "test_config" + strconv.FormatInt(time.Now().Unix(), 10)

	firstInsert := service.AddConfig(name, "test_value1")
	secondInsert := service.AddConfig(name, "test_value2")

	// Only pass t into top-level Convey calls
	Convey("test add config should succeed", t, func() {
		So(firstInsert, ShouldEqual, 1)
		So(secondInsert, ShouldEqual, 2)
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
