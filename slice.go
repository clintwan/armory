package armory

import (
	"errors"
	"reflect"
)

type slice struct{}

// Slice Slice
var Slice *slice

/*
func (arr Slice) IndexOf(ele interface{}) int {
	r := -1
	for idx, val := range arr {
		if ele == val {
			r = idx
		}
	}
	return r
}
*/

// IndexOf IndexOf
func (s *slice) IndexOf(params ...interface{}) (int, error) {
	arr := reflect.ValueOf(params[0])
	v := reflect.ValueOf(params[1])
	var t = reflect.TypeOf(params[0]).Kind()

	if t != reflect.Slice && t != reflect.Array {
		return -1, errors.New("Type Error! First argument must be an array or a slice")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Kind() != v.Kind() {
			return -1, errors.New("Type Error! Second argument must matched any elements in first argument")
		}
		if arr.Index(i).Interface() == v.Interface() {
			return i, nil
		}
	}
	return -1, nil
}
