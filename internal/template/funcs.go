package template

import (
	"reflect"
	"strings"
)

func unexported(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}

func deref(v interface{}) interface{} {
	val := reflect.ValueOf(v)

	if val.Kind() != reflect.Ptr {
		return v
	}

	return val.Elem().Interface()
}
