package goutils

import (
	"errors"
	"reflect"
	"strconv"
)

// Parse data update p by def
func Parse(data []string, def map[int]string, p interface{}) (map[string]string, error) {
	val := reflect.ValueOf(p)
	kd := val.Kind()
	if kd != reflect.Ptr {
		return nil, errors.New("参数不是指针")
	}
	property := make(map[string]string)
	l := len(data)
	for k, v := range def {
		if k >= l {
			continue
		}
		f := val.Elem().FieldByName(v)
		switch f.Kind() {
		case reflect.String:
			f.SetString(data[k])
		case reflect.Int:
			i, err := strconv.Atoi(data[k])
			if err == nil {
				f.SetInt(int64(i))
			}
		case reflect.Invalid:
			property[v] = data[k]
		}
	}
	fields := val.Elem().NumField()
	for i := 0; i < fields; i++ {
		switch val.Elem().Field(i).Kind() {
		case reflect.Map:
			val.Elem().Field(i).Set(reflect.ValueOf(property))
		}
	}
	return property, nil
}
