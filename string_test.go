package goutils

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSplitAfter(t *testing.T) {
	Convey("SpliteAfter", t, func() {
		Convey("12345 2,4", func() {
			ret := SplitAfter("12345", "2", "4")
			So(len(ret), ShouldEqual, 3)
			So(ret[0], ShouldEqual, "12")
			So(ret[1], ShouldEqual, "34")
			So(ret[2], ShouldEqual, "5")
		})
		Convey("12345 2,3", func() {
			ret := SplitAfter("12345", "2", "3")
			So(len(ret), ShouldEqual, 3)
			So(ret[0], ShouldEqual, "12")
			So(ret[1], ShouldEqual, "3")
			So(ret[2], ShouldEqual, "45")
		})
		Convey("12345 2,5", func() {
			ret := SplitAfter("12345", "2", "5")
			So(len(ret), ShouldEqual, 2)
			So(ret[0], ShouldEqual, "12")
			So(ret[1], ShouldEqual, "345")
		})
	})
}
