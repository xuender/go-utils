package goutils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type user struct {
	len int
}

func (u *user) Less(b interface{}) bool {
	if b == nil {
		return false
	}
	return u.len > b.(*user).len
}
func TestResults(t *testing.T) {
	Convey("Results", t, func() {
		rs := NewResults(3)
		So(rs.Len, ShouldEqual, 0)
		rs.Add(&user{len: 5})
		So(rs.Len, ShouldEqual, 1)
		rs.Add(&user{len: 15})
		So(rs.Len, ShouldEqual, 2)
		rs.Add(&user{len: 6})
		So(rs.Len, ShouldEqual, 3)
		rs.Add(&user{len: 85})
		So(rs.Len, ShouldEqual, 3)
		rs.Add(&user{len: 1})
		So(rs.Len, ShouldEqual, 3)
		So(rs.Get(0).(*user).len, ShouldEqual, 85)
		So(rs.Get(2).(*user).len, ShouldEqual, 6)
	})
}
