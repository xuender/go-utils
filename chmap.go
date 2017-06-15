package goutils

const (
	routePut = iota
	routeRemove
	routeError
)

type CallBack struct {
	Key    interface{}
	Value  interface{}
	ChBack chan CallBack
	Route  int
}

type ChMap struct {
	data       map[interface{}]interface{}
	chCallBack chan CallBack
}

func NewChMap() ChMap {
	chMap := ChMap{
		data:       make(map[interface{}]interface{}),
		chCallBack: make(chan CallBack, 3),
	}
	go chMap.run()
	return chMap
}

func (chMap ChMap) run() {
	for {
		cb := <-chMap.chCallBack
		switch cb.Route {
		case routePut:
			chMap.data[cb.Key] = cb.Value
		case routeRemove:
			if _, ok := chMap.data[cb.Key]; ok {
				delete(chMap.data, cb.Key)
			}
		case routeError:
			close(chMap.chCallBack)
			return
		}
		cb.ChBack <- cb
	}
}

func (p ChMap) Put(key, value interface{}) {
	ch := make(chan CallBack, 1)
	defer close(ch)
	p.chCallBack <- CallBack{
		Key:    key,
		Value:  value,
		Route:  routePut,
		ChBack: ch,
	}
	<-ch
}

func (p ChMap) Remove(key interface{}) {
	ch := make(chan CallBack, 1)
	defer close(ch)
	p.chCallBack <- CallBack{
		Key:    key,
		Route:  routeRemove,
		ChBack: ch,
	}
	<-ch
}

func (p ChMap) Close() {
	p.chCallBack <- CallBack{
		Route: routeError,
	}
}

func (p ChMap) Count() int {
	return len(p.data)
}

func (p ChMap) Get(key interface{}) (interface{}, bool) {
	v, ok := p.data[key]
	return v, ok
}

func (p ChMap) Has(key interface{}) bool {
	_, ok := p.data[key]
	return ok
}

func (p ChMap) Keys() []interface{} {
	keys := make([]interface{}, len(p.data))
	i := 0
	for k := range p.data {
		keys[i] = k
		i += 1
	}
	return keys
}

func (p ChMap) Iterator(callBack func(k, v interface{})) {
	for k, v := range p.data {
		if p.Has(k) {
			callBack(k, v)
		}
	}
}
