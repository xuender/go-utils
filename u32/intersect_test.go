package u32

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsIntersect(t *testing.T) {
	Convey("IsInterSect", t, func() {
		Convey("Empty", func() {
			So(IsIntersect(&[]uint32{}, &[]uint32{1}), ShouldEqual, false)
			So(IsIntersect(&[]uint32{1}, &[]uint32{}), ShouldEqual, false)
		})
		Convey("Dislocation", func() {
			So(IsIntersect(&[]uint32{1, 2}, &[]uint32{3, 4}), ShouldEqual, false)
			So(IsIntersect(&[]uint32{3, 4}, &[]uint32{1, 2}), ShouldEqual, false)
		})
		Convey("true", func() {
			So(IsIntersect(&[]uint32{1, 2}, &[]uint32{2, 3}), ShouldEqual, true)
			So(IsIntersect(&[]uint32{3, 4, 8}, &[]uint32{4, 5, 6}), ShouldEqual, true)
		})
		Convey("false", func() {
			So(IsIntersect(&[]uint32{1, 20}, &[]uint32{2, 3}), ShouldEqual, false)
			So(IsIntersect(&[]uint32{3, 4, 8}, &[]uint32{1, 50, 60}), ShouldEqual, false)
		})
	})
}
