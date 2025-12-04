package utils

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(label string, v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("%s: error = %v\n", label, err)
		return
	}
	fmt.Printf("%s:\n%s\n", label, string(b))
}
