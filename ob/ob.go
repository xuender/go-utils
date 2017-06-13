package ob

import (
	"github.com/xuender/goutils"
)

type Ob struct {
	ChMap   map[uint32]chan interface{}
	ChSuck  chan Suck
	MakeNum uint32
}

func NewOb(makeFunc func(ob *Ob)) *Ob {
	ob := &Ob{
		ChMap:   make(map[uint32]chan interface{}),
		ChSuck:  make(chan Suck, 1),
		MakeNum: 0,
	}
	go func() {
		for {
			suck := <-ob.ChSuck
			if suck.Close {
				delete(ob.ChMap, suck.Id)
			} else {
				ob.ChMap[suck.Id] = suck.ChData
			}
			for len(ob.ChMap) > 0 {
				makeFunc(ob)
				ob.MakeNum += 1
			}
		}
	}()
	return ob
}

func (ob *Ob) NewSuck() Suck {
	ch := make(chan interface{}, 1)
	e := Suck{Id: goutils.UniqueUint32(), ChData: ch}
	ob.ChSuck <- e
	return e
}

func (ob *Ob) Close(suck Suck) {
	suck.Close = true
	ob.ChSuck <- suck
}

func (ob *Ob) Notify(data interface{}) bool {
	select {
	case suck := <-ob.ChSuck:
		if suck.Close {
			delete(ob.ChMap, suck.Id)
		} else {
			ob.ChMap[suck.Id] = suck.ChData
		}
		return ob.update(data)
	default:
		return ob.update(data)
	}
}

func (ob *Ob) update(data interface{}) bool {
	if len(ob.ChMap) == 0 {
		return false
	}
	for _, ch := range ob.ChMap {
		ch <- data
	}
	return true
}
