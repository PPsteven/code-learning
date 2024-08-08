package main

import (
	"encoding/json"
	"fmt"
)

func isValidJson(s string) bool {
	return json.Valid([]byte(s))
}

func main() {
	jsons := []string{
		`{"a":1, "b":{"c":"123"}}`,
		`{"a":1`,
	}
	for _, jsonStr := range jsons {
		fmt.Printf(" %s ==> %v\n", jsonStr, isValidJson(jsonStr))
	}
	//{"a":1, "b":{"c":"123"}} ==> true
	//{"a":1 ==> false
}