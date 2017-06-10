package u32

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSort(t *testing.T) {
	Convey("Sort", t, func() {
		numbers := Int32Slice{14, 5, 7}
		numbers.Sort()
		So(numbers, ShouldResemble, Int32Slice{5, 7, 14})
	})
}
