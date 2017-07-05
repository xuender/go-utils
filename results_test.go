package goutils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestResults(t *testing.T) {
	Convey("Results", t, func() {
		rs := NewResults(3, func(i, j interface{}) bool {
			return i.(int) > j.(int)
		})
		So(rs.Len, ShouldEqual, 0)
		rs.Add(1, 5)
		So(rs.Len, ShouldEqual, 1)
		rs.Add(2, 15)
		So(rs.Len, ShouldEqual, 2)
		rs.Add(3, 6)
		So(rs.Len, ShouldEqual, 3)
		rs.Add(4, 85)
		So(rs.Len, ShouldEqual, 3)
		rs.Add(5, 1)
		So(rs.Len, ShouldEqual, 3)
		a, b := rs.Get(0)
		So(a, ShouldEqual, 4)
		So(b, ShouldEqual, 85)
		a, b = rs.Get(2)
		So(a, ShouldEqual, 3)
		So(b, ShouldEqual, 6)
		Convey("AddResults", func() {
			rs2 := NewResults(3, func(i, j interface{}) bool {
				return i.(int) > j.(int)
			})
			rs2.Add(7, 99)
			rs2.Add(9, 9)
			rs.AddResults(rs2)
			So(rs.GetData(0), ShouldEqual, 7)
			So(rs.GetPoint(0), ShouldEqual, 99)
		})
	})
}
