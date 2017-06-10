package goutils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestContains(t *testing.T) {
	Convey("Slice Contains", t, func() {
		x := []string{"aa", "bb", "cc"}
		Convey("Contains", func() {
			So(Contains(&x, func(i interface{}) bool {
				return i == "aa"
			}), ShouldEqual, true)
		})
		Convey("No Contains", func() {
			So(Contains(&x, func(i interface{}) bool {
				return i == "dd"
			}), ShouldEqual, false)
		})
	})
}
func TestRemove(t *testing.T) {
	Convey("Remove", t, func() {
		s := []int{1, 2, 4, 5}
		Convey("success", func() {
			Remove(&s, func(i interface{}) bool {
				return i == 4
			})
			So(len(s), ShouldEqual, 3)
			So(s[2], ShouldEqual, 5)
			Remove(&s, 1)
			So(len(s), ShouldEqual, 2)
			So(s[1], ShouldEqual, 5)
			Remove(&s, func(i int) bool {
				return i < 10
			})
			So(len(s), ShouldEqual, 0)
		})
		Convey("fail", func() {
			Remove(&s, 7)
			Remove(&s, func(i int) bool {
				return i > 100
			})
			So(len(s), ShouldEqual, 4)
		})
	})
}

/*
func TestFilter(t *testing.T) {
	Convey("Filter", t, func() {
		s := []int{1, 2, 4, 5}
		d := []int{1}
		Convey("success", func() {
			Filter(&s, &d, func(i interface{}) bool {
				return i == 4
			})
			So(len(d), ShouldEqual, 1)
			Filter(&s, &d, func(i int) bool {
				return i%2 == 0
			})
			So(len(d), ShouldEqual, 2)
		})
		Convey("fail", func() {
		})
	})
}
*/
