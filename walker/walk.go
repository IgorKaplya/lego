package walker

import (
	"fmt"
	"reflect"
)

func walk(x any, fn func(input string)) {
	val := getReflectValue(x)
	fmt.Printf("val is %#v\n", val.Interface())

	switch val.Kind() {
	case reflect.String:
		fmt.Printf("fn(%q)\n", val.String())
		fn(val.String())
	case reflect.Struct:
		fmt.Print("struct ->\n")
		for i := range val.NumField() {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		fmt.Print("slice ->\n")
		for i := range val.Len() {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.Map:
		fmt.Print("map ->\n")
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		fmt.Print("chan ->\n")
		for chanVal, ok := val.Recv(); ok; chanVal, ok = val.Recv() {
			walk(chanVal.Interface(), fn)
		}
	case reflect.Func:
		fmt.Print("func ->\n")
		for _, funcResult := range val.Call(nil) {
			walk(funcResult.Interface(), fn)
		}
	}
}

func getReflectValue(x any) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
