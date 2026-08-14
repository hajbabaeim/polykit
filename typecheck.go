package polykit

import (
	"fmt"
	"reflect"
)

func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Func, reflect.Slice, reflect.Interface:
		return value.IsNil()
	default:
		return false
	}
}

func IsEmpty(v interface{}) bool {
	if v == nil {
		return true
	}
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Array, reflect.String, reflect.Slice, reflect.Map:
		return value.Len() == 0
	case reflect.Ptr:
		if IsNil(v) {
			return true
		}
		return IsEmpty(value.Elem().Interface())
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			if !IsEmpty(value.Field(i).Interface()) {
				return false
			}
		}
	}
	return false
}

func ErrIfNotNil(v interface{}, message string) error {
	if !IsNil(v) {
		return fmt.Errorf("invalid value: %s, error: %v", v, message)
	}
	return nil
}
