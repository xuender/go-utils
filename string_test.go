package utils

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
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
		Convey("123", func() {
			ret := SplitAfter("123")
			So(len(ret), ShouldEqual, 1)
			So(ret[0], ShouldEqual, "123")
		})
	})
}

func TestIncludes(t *testing.T) {
	Convey("Includes", t, func() {
		Convey("包含", func() {
			array := []string{"1", "a"}
			So(Includes(array, "1"), ShouldEqual, true)
			So(Includes(array, "a"), ShouldEqual, true)
			So(Includes(array, "b"), ShouldEqual, false)
		})
	})
}

func ExampleSplitAfter() {
	fmt.Println(SplitAfter("110120119129", "0", "9"))

	// Output:
	// [110 120 119 129]
}
