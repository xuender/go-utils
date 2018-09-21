package goutils

import (
	"errors"
	"reflect"
)

// Filter 过滤
func Filter(collection, predicate interface{}) (interface{}, error) {
	rv := reflect.ValueOf(collection)
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return nil, errors.New("The passed collection is not a slice or an array")
	}
	fn := reflect.ValueOf(predicate)
	t := rv.Type().Elem()
	if !verifyFilterFuncType(fn, t) {
		return nil, errors.New("Function must be of type func(" + t.String() + ") bool or func(interface{}) bool")
	}
	var param [1]reflect.Value
	out := reflect.MakeSlice(rv.Type(), 0, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		param[0] = rv.Index(i)

		if fn.Call(param[:])[0].Bool() {
			out = reflect.Append(out, rv.Index(i))
		}

	}
	return out.Interface(), nil
}
func verifyFilterFuncType(fn reflect.Value, elType reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}
	if fn.Type().NumIn() != 1 || fn.Type().NumOut() != 1 {
		return false
	}
	return (fn.Type().In(0).Kind() == reflect.Interface || fn.Type().In(0) == elType) &&
		fn.Type().Out(0).Kind() == reflect.Bool
}
