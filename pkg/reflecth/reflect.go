package reflecth

import (
	"fmt"
	"reflect"
)

func GetName(f interface{}) string {
	val := reflect.TypeOf(f).Elem()
	return val.Name()
}

func GetStructValeByField(f interface{}, name string) interface{} {
	val := reflect.ValueOf(f).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		//tag := typeField.Tag

		if typeField.Name == name {
			fmt.Printf("%#v ", valueField)
			return valueField.Interface()
		}
	}

	return nil
}

func SetStructValeByField(f interface{}, fieldName string, value interface{}) error {
	val := reflect.ValueOf(f).Elem()
	if !val.CanAddr() {
		return fmt.Errorf("cannot assign to the item passed, item must be a pointer in order to assign")
	}

	findJsonName := map[string]reflect.Value{}

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		//tag := typeField.Tag

		findJsonName[typeField.Name] = valueField
	}

	fieldVal, ok := findJsonName[fieldName]
	if !ok {
		return fmt.Errorf("field %s does not exist within the provided item", fieldName)
	}

	fieldVal.Set(reflect.ValueOf(value))
	return nil
}
