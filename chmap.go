package goutils

const (
	typeGet = iota
	typeSet
	typeDel
	typeHas
	typeLen
	typeErr
)

type CallBack struct {
	Key    interface{}
	Value  interface{}
	ChBack chan CallBack
	Type   int
}

type ChMap chan CallBack

func NewChMap() ChMap {
	chMap := make(ChMap, 3)
	go func() {
		data := make(map[interface{}]interface{})
		for {
			cb := <-chMap
			switch cb.Type {
			case typeGet:
				if d, ok := data[cb.Key]; ok {
					cb.Value = d
				} else {
					cb.Type = typeErr
				}
				cb.ChBack <- cb
			case typeSet:
				data[cb.Key] = cb.Value
			case typeHas:
				_, ok := data[cb.Key]
				cb.Value = ok
				cb.ChBack <- cb
			case typeLen:
				cb.Value = len(data)
				cb.ChBack <- cb
			case typeDel:
				if _, ok := data[cb.Key]; ok {
					delete(data, cb.Key)
				}
			case typeErr:
				return
			}
		}
	}()
	return chMap
}

func (p ChMap) Get(key interface{}) (interface{}, bool) {
	ch := make(chan CallBack, 1)
	p <- CallBack{
		Key:    key,
		ChBack: ch,
		Type:   typeGet,
	}
	cb := <-ch
	if cb.Type == typeGet {
		return cb.Value, true
	}
	return nil, false
}

func (p ChMap) Set(key, value interface{}) {
	p <- CallBack{
		Key:   key,
		Value: value,
		Type:  typeSet,
	}
}

func (p ChMap) Has(key interface{}) bool {
	ch := make(chan CallBack, 1)
	p <- CallBack{
		Key:    key,
		ChBack: ch,
		Type:   typeHas,
	}
	cb := <-ch
	return cb.Value.(bool)
}

func (p ChMap) Len() int {
	ch := make(chan CallBack, 1)
	p <- CallBack{
		ChBack: ch,
		Type:   typeLen,
	}
	cb := <-ch
	return cb.Value.(int)
}

func (p ChMap) Del(key interface{}) {
	p <- CallBack{
		Key:  key,
		Type: typeDel,
	}
}

func (p ChMap) Close() {
	p <- CallBack{
		Type: typeErr,
	}
}
