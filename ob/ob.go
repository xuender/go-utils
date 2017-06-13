package ob

import (
	"github.com/xuender/goutils"
)

type Ob struct {
	ChMap     map[uint32]chan interface{}
	ChSuck    chan Suck
	DataMaker Maker
}

func NewOb(maker Maker) *Ob {
	ret := &Ob{
		DataMaker: maker,
		ChSuck:    make(chan Suck),
		ChMap:     make(map[uint32]chan interface{}),
	}
	go ret.run()
	return ret
}

func (s *Ob) run() {
	for {
		suck := <-s.ChSuck
		if suck.Remove {
			delete(s.ChMap, suck.Id)
		} else {
			s.ChMap[suck.Id] = suck.ChData
			s.DataMaker.Make(s)
		}
	}
}

func (s *Ob) NewSuck() Suck {
	ch := make(chan interface{})
	e := Suck{Id: goutils.UniqueUint32(), ChData: ch}
	s.ChSuck <- e
	return e
}

func (s *Ob) Close(suck Suck) {
	suck.Remove = true
	s.ChSuck <- suck
}

func (s *Ob) Notify(data interface{}) bool {
	select {
	case suck := <-s.ChSuck:
		if suck.Remove {
			delete(s.ChMap, suck.Id)
		} else {
			s.ChMap[suck.Id] = suck.ChData
		}
		return s.update(data)
	default:
		return s.update(data)
	}
}

func (s *Ob) update(data interface{}) bool {
	if len(s.ChMap) == 0 {
		return false
	}
	for _, ch := range s.ChMap {
		ch <- data
	}
	return true
}
