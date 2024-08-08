package arr_utils

import (
	"fmt"
	"reflect"
)

// PopArrByIndex implements go pop by reflect
func PopArrByIndex(arr interface{}, index int) interface{} {
	v := reflect.ValueOf(arr)
	k := reflect.TypeOf(arr)
	if v.Kind() != reflect.Ptr {
		panic(fmt.Errorf("%s should be ptr", k.Name()))
	}
	v = v.Elem()
	switch v.Kind() {
	case reflect.Array:
	case reflect.Slice:
		retVal := v.Index(index).Interface()
		v.Set(reflect.AppendSlice(v.Slice(0, index), v.Slice(index+1, v.Len())))
		return retVal
	default:
		panic(fmt.Errorf("elem of %s should be array or slice", k.Name()))
	}
	return nil
}
