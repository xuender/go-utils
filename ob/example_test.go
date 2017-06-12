package ob

import (
	"fmt"
	"sort"
	"time"
)

type TestMaker struct {
	Max int
}

func (t *TestMaker) Make(o *Ob) {
	for i := 0; i < t.Max; i++ {
		if !o.Notify(i) {
			break
		}
	}
}

func ExampleNewOb() {
	tm := new(TestMaker)
	tm.Max = 20
	o := NewOb(tm)
	ret := make([]int, 0)
	go func() {
		e := o.Add()
		for i := 0; i < 10; i++ {
			ret = append(ret, (<-e.ChOut).(int))
		}
		o.Remove(e)
	}()
	go func() {
		e := o.Add()
		for i := 0; i < 10; i++ {
			ret = append(ret, (<-e.ChOut).(int))
		}
		o.Remove(e)
	}()
	time.Sleep(5e9)
	sort.Ints(ret)
	fmt.Println(ret)
	// Output:
	// [0 0 1 1 2 2 3 3 4 4 5 5 6 6 7 7 8 8 9 9]
}
