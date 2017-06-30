package u32

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStrToUint32(t *testing.T) {
	Convey("string to uint32", t, func() {
		for i := 4294967200; i < 4294967295; i++ {
			So(StrToUint32(Uint32ToStr(uint32(i))), ShouldEqual, i)
		}
	})
}
