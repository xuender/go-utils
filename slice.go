package goutils

import (
	"reflect"
)

// Element contains by lambda or element.
func Contains(slice, lambdaOrElement interface{}) bool {
	sv := reflect.ValueOf(slice).Elem()
	if sv.Len() == 0 {
		return false
	}
	lv := reflect.ValueOf(lambdaOrElement)
	isFunc := lv.Kind() == reflect.Func
	for i := 0; i < sv.Len(); i++ {
		if isFunc {
			if lv.Call([]reflect.Value{sv.Index(i)})[0].Bool() {
				return true
			}
		} else {
			if reflect.DeepEqual(sv.Index(i).Interface(), lambdaOrElement) {
				return true
			}
		}
	}
	return false
}
// Remove element by lambda or element
func Remove(slice, lambdaOrElement interface{}) int {
	sv := reflect.ValueOf(slice).Elem()
	ret := 0
	if sv.Len() == 0 {
		return 0
	}
	lv := reflect.ValueOf(lambdaOrElement)
	isFunc := lv.Kind() == reflect.Func
	for i := 0; i < sv.Len(); i++ {
		isRemove := false
		if isFunc {
			isRemove = lv.Call([]reflect.Value{sv.Index(i)})[0].Bool()
		} else {
			isRemove = reflect.DeepEqual(sv.Index(i).Interface(), lambdaOrElement)
		}
		if isRemove {
			before := sv.Slice(0, i)
			after := sv.Slice(i + 1, sv.Len())
			reflect.Copy(sv, reflect.AppendSlice(before, after))
			sv.SetLen(sv.Len() - 1)
			ret += 1
			i = -1
		}
	}
	return ret
}

