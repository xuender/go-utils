package ob

import (
	"fmt"
	"sort"
	"time"
)

type testMaker struct {
}

func (t *testMaker) Make(o *Ob) {
	i := 0
	for {
		if !o.Notify(i) {
			return
		}
		i++
	}
}

func ExampleNewOb() {
	tm := new(testMaker)
	ob := NewOb(tm)
	ret := make([]int, 0)
	// for f := 0; f < 2; f++ {
	go func() {
		suck := ob.NewSuck()
		defer ob.Close(suck)
		for i := 0; i < 10; i++ {
			ret = append(ret, (<-suck.ChData).(int))
		}
	}()
	// }
	sort.Ints(ret)
	time.Sleep(1e9)
	fmt.Println(ret)
	// Output:
	// [0 1 2 3 4 5 6 7 8 9]
}
