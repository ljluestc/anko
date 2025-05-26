package vm

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// toString converts all reflect.Value-s into string.
func toString(v reflect.Value) string {
	if v.Kind() == reflect.Interface && !v.IsNil() {
		v = v.Elem()
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == reflect.String {
		return v.String()
	}
	return fmt.Sprint(v.Interface())
}

// toBool converts all reflect.Value-s into bool.
func toBool(v reflect.Value) bool {
	b, _ := tryToBool(v)
	return b
}

// toFloat64 converts all reflect.Value-s into float64.
func toFloat64(v reflect.Value) float64 {
	f, _ := tryToFloat64(v)
	return f
}

// toInt64 converts all reflect.Value-s into int64.
func toInt64(v reflect.Value) int64 {
	i, _ := tryToInt64(v)
	return i
}

// toInt converts all reflect.Value-s into int.
func toInt(v reflect.Value) int {
	i, _ := tryToInt(v)
	return i
}

// tryToInt attempts to convert a value to an int.
// If it cannot (in the case of a non-numeric string, a struct, etc.)
// it returns 0 and an error.
func tryToInt(v reflect.Value) (int, error) {
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Float64, reflect.Float32:
		return int(v.Float()), nil
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		return int(v.Int()), nil
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	case reflect.String:
		s := v.String()
		var i int64
		var err error
		if strings.HasPrefix(s, "0x") {
			i, err = strconv.ParseInt(s, 16, 64)
		} else {
			i, err = strconv.ParseInt(s, 10, 64)
		}
		if err == nil {
			return int(i), nil
		}
	}
	return 0, errors.New("couldn't convert to integer")
}
