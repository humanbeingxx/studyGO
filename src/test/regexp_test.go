package test

import (
	"regexp"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

var digital, _ = regexp.Compile("^\\d+$")

func TestDigital(t *testing.T) {
	convey.Convey("is digital", t, func() {
		convey.So(digital.Match([]byte("12345")), convey.ShouldEqual, true)
	})

	convey.Convey("is not digital", t, func() {
		convey.So(digital.Match([]byte("12345xx")), convey.ShouldNotEqual, true)
	})
}
