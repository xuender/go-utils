package goutils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFilter(t *testing.T) {
	Convey("Filter", t, func() {
		arr := [4]int{1, 2, 3, 4}
		Convey("Array", func() {
			na, err := Filter(arr, func(i int) bool {
				return i == 3
			})
			So(err, ShouldBeNil)
			So(len(arr), ShouldEqual, 4)
			a, ok := na.([]int)
			So(ok, ShouldEqual, true)
			if ok {
				So(len(a), ShouldEqual, 1)
			}
		})
		Convey("Slice", func() {
			ss := arr[:]
			Filter(&ss, func(i int) bool {
				return i == 3
			})
			So(len(ss), ShouldEqual, 1)
		})
		Convey("Slice 2", func() {
			ss := arr[:]
			na, _ := Filter(ss, func(i int) bool {
				return i == 3
			})
			a, ok := na.([]int)
			So(ok, ShouldEqual, true)
			if ok {
				So(len(a), ShouldEqual, 1)
			}
			So(len(ss), ShouldEqual, 4)
		})
	})
}
