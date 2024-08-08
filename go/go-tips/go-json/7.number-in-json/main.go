package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// handle numbers in json: json.Decoder and json.Decoder.UseNumber
func main() {
	s := `{"float":3.12, "int":3.0}`
	var res interface{}
	_ = json.Unmarshal([]byte(s), &res)
	fmt.Printf("%v\n", res)

	// map[float:3.12 int:3]
	// 3.0 was transfer to 3

	d := json.NewDecoder(bytes.NewReader([]byte(s)))
	// UseNumber 直接存储的是字符串，所以可以保持3.0 不变
	d.UseNumber()
	_ = d.Decode(&res)
	fmt.Printf("%v\n", res)
	// map[float:3.12 int:3.0]
}