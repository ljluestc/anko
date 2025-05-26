package vm

import (
	"reflect"
)

// makeValue creates a new reflect.Value for the given type.
func makeValue(t reflect.Type) (reflect.Value, error) {
	if t.Kind() == reflect.Ptr {
		return reflect.New(t.Elem()), nil
	}
	return reflect.Zero(t), nil
}

// falseValue and trueValue represent boolean constants.
var (
	falseValue = reflect.ValueOf(false)
	trueValue  = reflect.ValueOf(true)
)

// getMapIndex retrieves the value for a key in a map.
func getMapIndex(key, m reflect.Value) reflect.Value {
	if !key.IsValid() || !m.IsValid() || m.Kind() != reflect.Map {
		return reflect.Value{}
	}
	return m.MapIndex(key)
}

// stringType represents the reflect.Type for a string.
var stringType = reflect.TypeOf("")

// isNil checks if a reflect.Value is nil.
func isNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}
