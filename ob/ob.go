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
		ChSuck:    make(chan Suck, 1),
		ChMap:     make(map[uint32]chan interface{}),
	}
	go ret.run()
	return ret
}

func (ob *Ob) run() {
	for {
		suck := <-ob.ChSuck
		if suck.Close {
			delete(ob.ChMap, suck.Id)
		} else {
			ob.ChMap[suck.Id] = suck.ChData
		}
		if len(ob.ChMap) > 0 {
			ob.DataMaker.Make(ob)
		}
	}
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
