package utils

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/goccy/go-yaml"
)

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

func YamlStringToJSON(yamlStr string) ([]byte, error) {
	var raw interface{}

	// Parse YAML string
	if err := yaml.Unmarshal([]byte(yamlStr), &raw); err != nil {
		return nil, err
	}

	raw = convertKeysToString(raw)

	jsonData, err := MapToJSON(raw)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func convertKeysToString(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m := make(map[string]interface{})
		for k, v := range x {
			m[fmt.Sprint(k)] = convertKeysToString(v)
		}
		return m

	case map[string]interface{}:
		m := make(map[string]interface{})
		for k, v := range x {
			m[k] = convertKeysToString(v)
		}
		return m

	case []interface{}:
		for i, v := range x {
			x[i] = convertKeysToString(v)
		}
		return x
	}

	return i
}

func JSONStringToMap(jsonStr string) (map[string]interface{}, error) {
	var result map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func MapToJSON(m interface{}) ([]byte, error) {
	return json.MarshalIndent(m, "", "  ")
}
