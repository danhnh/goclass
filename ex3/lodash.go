package ex3

import (
	"errors"
	"reflect"
)

func IsEmpty(v interface{}) bool {
	if reflect.TypeOf(v) == nil {
		return true
	}
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int, reflect.Int64, reflect.Int8, reflect.Int16, reflect.Int32:
		return reflect.ValueOf(v).Int() == 0
	case reflect.Uint, reflect.Uint64, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return reflect.ValueOf(v).Uint() == 0
	case reflect.Float32, reflect.Float64:
		return reflect.ValueOf(v).Float() == 0
	case reflect.Bool:
		return !reflect.ValueOf(v).Bool()
	case reflect.Slice, reflect.Array, reflect.Chan, reflect.Map, reflect.String:
		return reflect.ValueOf(v).Len() == 0
	case reflect.Ptr, reflect.UnsafePointer:
		return reflect.ValueOf(v).Pointer() == 0
	}
	return false
}

func getItems(v interface{}) []reflect.Value {
	var result []reflect.Value
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice, reflect.Array:
		len := reflect.ValueOf(v).Len()
		for i := 0; i < len; i++ {
			result = append(result, reflect.ValueOf(v).Index(i))
		}
	//case reflect.Map:
	//	len := reflect.ValueOf(v).Len()
	//	keys := reflect.ValueOf(v).MapKeys()
	//	for i := 0; i < len; i++ {
	//		key := keys[i]
	//		result = append(result, reflect.ValueOf(v).MapIndex(key))
	//	}
	}
	return result
}

func Last(v interface{}) (interface{}, error) {
	if reflect.TypeOf(v) == nil {
		return nil, errors.New("parameter cannot be nil")
	}
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice, reflect.Array:
		len := reflect.ValueOf(v).Len()
		if len == 0 {
			return nil, errors.New(reflect.TypeOf(v).Kind().String() + "is empty")
		}
		return reflect.ValueOf(v).Index(len - 1).Interface(), nil
	}
	return nil, errors.New("parameter is not valid")
}

func Map(v interface{}, f interface{}) interface{} {
	if reflect.TypeOf(v) == nil || (reflect.TypeOf(v).Kind() != reflect.Array && reflect.TypeOf(v).Kind() != reflect.Slice) {
		panic("parameter v must be map, array, slice")
	}
	if reflect.TypeOf(f) == nil || reflect.TypeOf(f).Kind() != reflect.Func || reflect.ValueOf(f).Type().NumOut() != 1 {
		panic("parameter f must be func with 1 output value")
	}
	var result reflect.Value
	if reflect.TypeOf(f).Kind() == reflect.Func {
		result = reflect.MakeSlice(reflect.SliceOf(reflect.ValueOf(f).Type().Out(0)), 0, 0)
	}
	vVal := getItems(v)
	if len(vVal) == 0 {
		return result.Interface()
	}
	//switch reflect.TypeOf(f).Kind() {
	//case reflect.String:
	//	for i := 0; i < len(vVal); i++ {
	//		if vVal[i].Kind() == reflect.Struct {
	//			result = reflect.Append(result, vVal[i].FieldByName(reflect.ValueOf(f).String()))
	//		} else if vVal[i].Kind() == reflect.Map {
	//			result = reflect.Append(result, vVal[i].MapIndex(reflect.ValueOf(f)))
	//		}
	//	}
	//case reflect.Func:
		for i := 0; i < len(vVal); i++ {
			var params []reflect.Value
			for z:=0;z<reflect.ValueOf(f).Type().NumIn();z++ {
				params = append(params, vVal[i])
			}
			result = reflect.Append(result, reflect.ValueOf(f).Call(params)[0])
		}
	//}
	return result.Interface()
}

func isGreater(v1, v2 interface{}) (bool, error) {
	switch reflect.TypeOf(v1).Kind() {
	case reflect.Int, reflect.Int64, reflect.Int8, reflect.Int16, reflect.Int32:
		return reflect.ValueOf(v1).Int() > reflect.ValueOf(v2).Int(), nil
	case reflect.Uint, reflect.Uint64, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return reflect.ValueOf(v1).Uint() > reflect.ValueOf(v2).Uint(), nil
	case reflect.Float32, reflect.Float64:
		return reflect.ValueOf(v1).Float() > reflect.ValueOf(v2).Float(), nil
	}
	return false, errors.New("Cannot compare " + reflect.TypeOf(v1).Kind().String())
}

func Max(v interface{}) (interface{}, error) {
	if reflect.TypeOf(v) == nil {
		return nil, errors.New("parameter cannot be nil")
	}
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice, reflect.Array:
		len := reflect.ValueOf(v).Len()
		if len == 0 {
			return nil, errors.New(reflect.TypeOf(v).Kind().String() + "is empty")
		}
		max := reflect.ValueOf(v).Index(0).Interface()
		for i := 1; i < len; i++ {
			val := reflect.ValueOf(v).Index(i).Interface()
			b, e := isGreater(val, max)
			if e != nil {
				return nil, e
			}
			if b {
				max = val
			}
		}
		return max, nil
	}
	return nil, errors.New("parameter is not valid")
}

func IndexOf(v1, v2 interface{}, fromIndex int) int {
	if reflect.TypeOf(v1) == nil || !reflect.TypeOf(v2).Comparable() {
		return -1
	}
	switch reflect.TypeOf(v1).Kind() {
	case reflect.Slice, reflect.Array:
		len := reflect.ValueOf(v1).Len()
		if len == 0 || fromIndex >= len || fromIndex < 0 {
			return -1
		}
		for i := fromIndex; i < len; i++ {
			if reflect.ValueOf(v1).Index(i).Interface() == reflect.ValueOf(v2).Interface() {
				return i
			}
		}
	}
	return -1
}

// func Head(v interface{}) interface{} {

// 	return nil
// }
