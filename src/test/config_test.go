package test

import (
	"fmt"
<<<<<<< HEAD
	"math/rand"
	"service"
	"strconv"
	"testing"
	"time"
=======
	"service"
	"testing"
>>>>>>> 4615ee5c1f671a4b947a970b5d3c0f76bbb9bef7

	. "github.com/smartystreets/goconvey/convey"
)

<<<<<<< HEAD
func TestRandInt(t *testing.T) {
	ran := rand.Intn(100)

	fmt.Println(ran)
}

func TestNow(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Unix())
}

=======
>>>>>>> 4615ee5c1f671a4b947a970b5d3c0f76bbb9bef7
func TestReadConfig(t *testing.T) {
	configs := service.ReadConfig()
	fmt.Print(configs)
}
<<<<<<< HEAD

func TestAddConfig(t *testing.T) {
	name := "test_config" + strconv.FormatInt(time.Now().Unix(), 10)

	firstInsert := service.AddConfig(name, "test_value1")
	secondInsert := service.AddConfig(name, "test_value2")

	// Only pass t into top-level Convey calls
	Convey("test add config should succeed", t, func() {
		So(firstInsert, ShouldEqual, 1)
		So(secondInsert, ShouldEqual, 2)
=======
func TestAddConfig(t *testing.T) {
	success := service.AddConfig("test_confi1", "test_confi2")

	// Only pass t into top-level Convey calls
	Convey("test add config should succeed", t, func() {
		So(success, ShouldEqual, true)
>>>>>>> 4615ee5c1f671a4b947a970b5d3c0f76bbb9bef7
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
