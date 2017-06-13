package ob

import (
	"fmt"
	"time"
)

type testMaker struct {
	Max int
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
	tm.Max = 20
	o := NewOb(tm)
	ret := make([]int, 0)
	// for i := 0; i < 2; i++ {
	go func() {
		e := o.NewSuck()
		defer o.Close(e)
		for i := 0; i < 10; i++ {
			ret = append(ret, (<-e.ChData).(int))
		}
	}()
	// }
	time.Sleep(1e9)
	fmt.Println(ret)
	// Output:
	// [0 1 2 3 4 5 6 7 8 9]
}
