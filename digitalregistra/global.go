package digitalregistra

import (
	"fmt"
	"net/url"
	"reflect"
)

func mapToURLValues(inter map[string]interface{}) url.Values {
	data := url.Values{}

	for name, val := range inter {
		if reflect.ValueOf(val).Kind() == reflect.Ptr && reflect.ValueOf(val).IsNil() {
			continue
		}
		if reflect.ValueOf(val).Kind() == reflect.Ptr && fmt.Sprintf("%v", reflect.Indirect(reflect.ValueOf(val))) != "" {
			data.Add(name, fmt.Sprintf("%v", reflect.Indirect(reflect.ValueOf(val))))
		} else if reflect.ValueOf(val).Kind() != reflect.Ptr && fmt.Sprintf("%v", val) != "" {
			data.Add(name, fmt.Sprintf("%v", val))
		}
	}

	return data
}

func ModelToURLValues(inter interface{}) url.Values {
	data := url.Values{}

	rfl := reflect.ValueOf(inter)
	typ := reflect.TypeOf(inter)

	fieldLen := typ.NumField()

	for i := 0; i < fieldLen; i++ {
		if rfl.Field(i).Kind() == reflect.Ptr {
			if rfl.Field(i).IsNil() {
				continue
			}
			data.Add(typ.Field(i).Tag.Get("digitalregistra"), fmt.Sprintf("%+v", reflect.Indirect(rfl.Field(i))))
		} else {
			data.Add(typ.Field(i).Tag.Get("digitalregistra"), fmt.Sprintf("%+v", rfl.Field(i).Interface()))
		}
	}
	return data
}
