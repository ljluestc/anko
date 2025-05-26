package vm

import (
    "fmt"
    "reflect"
)

// invokeMember handles member access (e.g., obj.member or obj.method()) in Anko.
func invokeMember(v reflect.Value, name string) (reflect.Value, error) {
    // Handle direct field access.
    if v.Kind() == reflect.Struct {
        if field := v.FieldByName(name); field.IsValid() {
            return field, nil
        }
    }

    // Handle method access.
    method := v.MethodByName(name)
    if method.IsValid() {
        return method, nil
    }

    // Fix: Try pointer receiver methods if v is a struct value.
    if v.Kind() == reflect.Struct && !v.CanAddr() {
        // Create a pointer to the struct value.
        ptr := reflect.New(v.Type())
        ptr.Elem().Set(v)
        method = ptr.MethodByName(name)
        if method.IsValid() {
            return method, nil
        }
    }

    // Try pointer receiver methods if v is addressable.
    if v.CanAddr() {
        method = v.Addr().MethodByName(name)
        if method.IsValid() {
            return method, nil
        }
    }

    return reflect.Value{}, fmt.Errorf("no member named '%s' for %s", name, v.Type())
}
