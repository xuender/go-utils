package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParse(t *testing.T) {
	Convey("Parse", t, func() {
		type ts struct {
			Name       string
			Num        int
			Properties map[string]string
		}
		Convey("p", func() {
			t := &ts{}
			Parse([]string{"张三", "25", "中国", "汉族"}, map[int]string{0: "Name", 1: "Num", 2: "国籍", 3: "民族"}, t)
			So(t.Name, ShouldEqual, "张三")
			So(t.Num, ShouldEqual, 25)
			So(len(t.Properties), ShouldEqual, 2)
			So(t.Properties["国籍"], ShouldEqual, "中国")
			So(t.Properties["民族"], ShouldEqual, "汉族")
		})
	})
}
