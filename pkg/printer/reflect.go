package printer

import (
	"errors"
	"reflect"
)

// FieldMap returns a map of all field names and values that can be used
// at runtime without panicking.
func FieldMap(anystruct any) (map[string]any, error) {
	val, typ := reflect.ValueOf(anystruct), reflect.TypeOf(anystruct)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	m := make(map[string]any)
	if val.Kind() != reflect.Struct {
		return m, errors.New("error: input must be a struct type")
	}
	for i := range val.NumField() {
		fieldVal := val.Field(i)
		if fieldVal.CanInterface() {
			m[typ.Field(i).Name] = fieldVal
		}
	}
	return m, nil
}

// FieldNames returns a string list of all field names that can be used
// at runtime without panicking.
func FieldNames(anystruct any) ([]string, error) {
	val, typ := reflect.ValueOf(anystruct), reflect.TypeOf(anystruct)
	if val.Kind() == reflect.Ptr { // dereference a pointer
		val = val.Elem()
		typ = typ.Elem()
	}
	names := []string{}
	if val.Kind() != reflect.Struct {
		return names, errors.New("error: input must be a struct type")
	}
	for i := range val.NumField() {
		if val.Field(i).CanInterface() {
			names = append(names, typ.Field(i).Name)
		}
	}
	return names, nil
}
