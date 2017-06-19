package goutils

const (
	routePut = iota
	routeRemove
	routeError
)

type callBack struct {
	Key    interface{}
	Value  interface{}
	ChBack chan callBack
	Route  int
}

// ChMap is channel map.
type ChMap struct {
	data       map[interface{}]interface{}
	chCallBack chan callBack
}

// Put value by key.
func (p ChMap) Put(key, value interface{}) {
	ch := make(chan callBack, 1)
	defer close(ch)
	p.chCallBack <- callBack{
		Key:    key,
		Value:  value,
		Route:  routePut,
		ChBack: ch,
	}
	<-ch
}

// Remove obj by key.
func (p ChMap) Remove(key interface{}) {
	ch := make(chan callBack, 1)
	defer close(ch)
	p.chCallBack <- callBack{
		Key:    key,
		Route:  routeRemove,
		ChBack: ch,
	}
	<-ch
}

// Close this ChMap.
func (p ChMap) Close() {
	p.chCallBack <- callBack{
		Route: routeError,
	}
}

// Count ChMap.
func (p ChMap) Count() int {
	return len(p.data)
}

// Get obj by key.
func (p ChMap) Get(key interface{}) (interface{}, bool) {
	v, ok := p.data[key]
	return v, ok
}

// Has key.
func (p ChMap) Has(key interface{}) bool {
	_, ok := p.data[key]
	return ok
}

// Keys is get this map keys.
func (p ChMap) Keys() []interface{} {
	keys := make([]interface{}, len(p.data))
	i := 0
	for k := range p.data {
		keys[i] = k
		i++
	}
	return keys
}

// Iterator map.
func (p ChMap) Iterator(callBack func(k, v interface{})) {
	for k, v := range p.data {
		if p.Has(k) {
			callBack(k, v)
		}
	}
}

func (p ChMap) run() {
	for {
		cb := <-p.chCallBack
		switch cb.Route {
		case routePut:
			p.data[cb.Key] = cb.Value
		case routeRemove:
			if _, ok := p.data[cb.Key]; ok {
				delete(p.data, cb.Key)
			}
		case routeError:
			close(p.chCallBack)
			return
		}
		cb.ChBack <- cb
	}
}

// NewChMap new ChMap.
func NewChMap() ChMap {
	chMap := ChMap{
		data:       make(map[interface{}]interface{}),
		chCallBack: make(chan callBack, 3),
	}
	go chMap.run()
	return chMap
}
