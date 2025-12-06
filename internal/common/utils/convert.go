package utils

import "reflect"

func TypeToJson(v interface{}) map[string]string {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	desc := map[string]string{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		desc[f.Name] = f.Tag.Get("desc")
	}
	return desc
}
