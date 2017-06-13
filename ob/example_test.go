package ob

import (
	"fmt"
	"sort"
	"time"
)

func ExampleNewOb() {
	ob := NewOb(func(ob *Ob) {
		i := 0
		for {
			if !ob.Notify(i) {
				return
			}
			i += 1
		}
	})
	ret := make([]int, 0)
	for f := 0; f < 2; f++ {
		go func() {
			suck := ob.NewSuck()
			defer ob.Close(suck)
			for i := 0; i < 3; i++ {
				ret = append(ret, (<-suck.ChData).(int))
			}
		}()
	}
	time.Sleep(1e9)
	sort.Ints(ret)
	fmt.Println(ret)
	// Output:
	// [0 0 1 1 2 2]
}
