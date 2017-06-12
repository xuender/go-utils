package ob

import (
	"github.com/xuender/goutils"
)

type Ob struct {
	Observers  map[uint32]chan interface{}
	ChEvent    chan Event
	Observable Maker
}

func NewOb(maker Maker) *Ob {
	ret := &Ob{
		Observable: maker,
		ChEvent:    make(chan Event),
		Observers:  make(map[uint32]chan interface{}),
	}
	go ret.run()
	return ret
}

func (s *Ob) Add() Event {
	ch := make(chan interface{})
	e := Event{Id: goutils.UniqueUint32(), ChOut: ch}
	s.ChEvent <- e
	return e
}

func (s *Ob) Remove(event Event) {
	event.Remove = true
	s.ChEvent <- event
}

func (s *Ob) run() {
	for {
		event := <-s.ChEvent
		if event.Remove {
			// fmt.Printf("删除:%d\n", event.Id)
			delete(s.Observers, event.Id)
		} else {
			// fmt.Printf("增加:%d\n", event.Id)
			s.Observers[event.Id] = event.ChOut
			s.Observable.Make(s)
		}
	}
}

func (s *Ob) Notify(data interface{}) bool {
	select {
	case event := <-s.ChEvent:
		if event.Remove {
			// fmt.Printf("删除:%d\n", event.Id)
			delete(s.Observers, event.Id)
		} else {
			// fmt.Printf("增加:%d\n", event.Id)
			s.Observers[event.Id] = event.ChOut
		}
		return s.update(data)
	default:
		return s.update(data)
	}
}

func (s *Ob) update(data interface{}) bool {
	if len(s.Observers) == 0 {
		return false
	}
	for _, ch := range s.Observers {
		ch <- data
	}
	return true
}
