package u32

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEncode(t *testing.T) {
	Convey("encode", t, func() {
		So(Decode(Encode(33)), ShouldEqual, 33)
	})
	Convey("nil", t, func() {
		So(Decode([]byte{}), ShouldEqual, 0)
	})
}
