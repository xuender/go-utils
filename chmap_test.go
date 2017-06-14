package goutils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCmap(t *testing.T) {
	Convey("Cmap", t, func() {
		chMap := NewChMap()
		chMap.Set("1", 1)
		chMap.Set("2", 2)
		Convey("Get", func() {
			v, ok := chMap.Get("2")
			So(v, ShouldEqual, 2)
			So(ok, ShouldEqual, true)
		})
		Convey("Set", func() {
			chMap.Set("2", 3)
			v, ok := chMap.Get("2")
			So(v, ShouldEqual, 3)
			So(ok, ShouldEqual, true)
		})
		Convey("Len", func() {
			So(chMap.Len(), ShouldEqual, 2)
		})
		Convey("Has", func() {
			So(chMap.Has("1"), ShouldEqual, true)
			So(chMap.Has("no"), ShouldEqual, false)
		})
		Convey("Del", func() {
			chMap.Del("1")
			So(chMap.Len(), ShouldEqual, 1)
		})
	})
}
