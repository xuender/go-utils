package goutils

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCache(t *testing.T) {
	Convey("Cache", t, func() {
		cache := NewCache(time.Second * 1)
		defer cache.Close()
		cache.Put(1, 1)
		cache.Put(2, 2)
		Convey("Count", func() {
			So(cache.Count(), ShouldEqual, 2)
		})
		Convey("Get", func() {
			v, ok := cache.Get(2)
			So(v, ShouldEqual, 2)
			So(ok, ShouldEqual, true)
		})
		/*
			Convey("Keys", func() {
				keys := chMap.Keys()
				So(len(keys), ShouldEqual, 2)
			})
		*/
		Convey("Put", func() {
			cache.Put(2, 3)
			v, ok := cache.Get(2)
			So(v, ShouldEqual, 3)
			So(ok, ShouldEqual, true)
		})
		Convey("Remove", func() {
			cache.Remove(1)
			So(cache.Count(), ShouldEqual, 1)
		})
		Convey("Time", func() {
			time.Sleep(time.Second * 2)
			So(cache.Count(), ShouldEqual, 0)
		})

	})
}
