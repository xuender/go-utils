package u32

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJaccard(t *testing.T) {
	Convey("Jaccard", t, func() {
		Convey("Empty", func() {
			So(Jaccard(&[]uint32{}, &[]uint32{1}), ShouldEqual, 0)
			So(Jaccard(&[]uint32{1}, &[]uint32{}), ShouldEqual, 0)
		})
		Convey("Dislocation", func() {
			So(Jaccard(&[]uint32{1, 2}, &[]uint32{3, 4}), ShouldEqual, 0)
			So(Jaccard(&[]uint32{3, 4}, &[]uint32{1, 2}), ShouldEqual, 0)
		})
		Convey("1000", func() {
			So(Jaccard(&[]uint32{1, 2}, &[]uint32{1, 2}), ShouldEqual, 1000)
		})
		Convey("Repeat", func() {
			So(Jaccard(&[]uint32{1, 2, 3}, &[]uint32{1, 2, 4, 1, 1, 4}), ShouldEqual, 500)
			So(Jaccard(&[]uint32{2, 3, 1, 2, 3}, &[]uint32{1, 2, 4}), ShouldEqual, 500)
		})
	})
}
