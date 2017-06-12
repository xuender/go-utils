package ob

import (
	"fmt"
	"time"

	"../goutils"
)

func main() {
	fmt.Println("start")
	chEvent := make(chan pool.Event)
	go data(chEvent)
	go run(chEvent)
	// time.Sleep(1e9)
	go run(chEvent)
	time.Sleep(3e9)
	fmt.Println("end")
}

func run(chEvent chan pool.Event) {
	ch := make(chan interface{})
	e := pool.Event{Id: goutils.UniqueUint32(), ChOut: ch}
	chEvent <- e
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	e.Remove = true
	chEvent <- e
}

func data(in chan pool.Event) {
	m := make(map[uint32]chan interface{})
	for {
		a := <-in
		if a.Remove {
			delete(m, a.Id)
		} else {
			fmt.Printf("增加:%d\n", a.Id)
			m[a.Id] = a.ChOut
			for i := 100; i < 10000; i++ {
				select {
				case b := <-in:
					if b.Remove {
						delete(m, b.Id)
					} else {
						fmt.Printf("增加:%d\n", b.Id)
						m[b.Id] = b.ChOut
						if len(m) > 0 {
							for _, item := range m {
								item <- i
							}
						}
					}
				default:
					if len(m) > 0 {
						for _, item := range m {
							item <- i
						}
					} else {
						break
					}
				}
			}
		}
	}
}
