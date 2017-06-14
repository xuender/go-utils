package goutils

const (
	routeGet = iota
	routeHas
	routeCount
	routePut
	routeKeys
	routeRemove
	routeError
)

type CallBack struct {
	Key    interface{}
	Value  interface{}
	ChBack chan CallBack
	Route  int
}

type ChMap chan CallBack

func NewChMap() ChMap {
	chMap := make(ChMap, 3)
	go func() {
		data := make(map[interface{}]interface{})
		for {
			cb := <-chMap
			switch cb.Route {
			case routeCount:
				cb.Value = len(data)
				cb.ChBack <- cb
			case routeGet:
				if d, ok := data[cb.Key]; ok {
					cb.Value = d
				} else {
					cb.Route = routeError
				}
				cb.ChBack <- cb
			case routeHas:
				_, ok := data[cb.Key]
				cb.Value = ok
				cb.ChBack <- cb
			case routeKeys:
				keys := make([]interface{}, len(data))
				i := 0
				for k := range data {
					keys[i] = k
					i += 1
				}
				cb.Value = keys
				cb.ChBack <- cb
			case routePut:
				data[cb.Key] = cb.Value
			case routeRemove:
				if _, ok := data[cb.Key]; ok {
					delete(data, cb.Key)
				}
			case routeError:
				close(chMap)
				return
			}
		}
	}()
	return chMap
}

func (p ChMap) Count() int {
	ch := make(chan CallBack, 1)
	defer close(ch)
	p <- CallBack{
		ChBack: ch,
		Route:  routeCount,
	}
	cb := <-ch
	return cb.Value.(int)
}

func (p ChMap) Get(key interface{}) (interface{}, bool) {
	ch := make(chan CallBack, 1)
	defer close(ch)
	p <- CallBack{
		Key:    key,
		ChBack: ch,
		Route:  routeGet,
	}
	cb := <-ch
	if cb.Route == routeGet {
		return cb.Value, true
	}
	return nil, false
}

func (p ChMap) Has(key interface{}) bool {
	ch := make(chan CallBack, 1)
	defer close(ch)
	p <- CallBack{
		Key:    key,
		ChBack: ch,
		Route:  routeHas,
	}
	cb := <-ch
	return cb.Value.(bool)
}

func (p ChMap) Keys() []interface{} {
	ch := make(chan CallBack, 1)
	defer close(ch)
	p <- CallBack{
		ChBack: ch,
		Route:  routeKeys,
	}
	cb := <-ch
	return cb.Value.([]interface{})
}

func (p ChMap) Put(key, value interface{}) {
	p <- CallBack{
		Key:   key,
		Value: value,
		Route: routePut,
	}
}

func (p ChMap) Remove(key interface{}) {
	p <- CallBack{
		Key:   key,
		Route: routeRemove,
	}
}

func (p ChMap) Close() {
	p <- CallBack{
		Route: routeError,
	}
}
