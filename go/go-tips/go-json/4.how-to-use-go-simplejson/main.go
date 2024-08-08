package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"strconv"
)

func main() {
	buf := bytes.NewBuffer([]byte(`{
		"test": {
			"array": [1, "2", 3],
			"arraywithsubs": [
				{"subkeyone": 1},
				{"subkeytwo": 2, "subkeythree": 3}
			],
			"bignum": 8000000000
		}
	}`))
	js, _ := simplejson.NewFromReader(buf)

	arr, _ := js.Get("test").Get("array").Array()
	fmt.Printf("arr: %v\n", arr)
	for i, v := range arr {
		var iv int
		switch v.(type) {
		case float64:
			iv = int(v.(float64))
		case string:
			iv, _ = strconv.Atoi(v.(string))
		case json.Number:
			i64, _ := v.(json.Number).Int64()
			iv = int(i64)
		}
		fmt.Printf("index[%d]: %v\n", i, iv)
	}

	ma := js.Get("test").Get("array").MustArray()
	fmt.Printf("arr: %v\n", ma)

	mm := js.Get("test").Get("arraywithsubs").GetIndex(0).MustMap()
	fmt.Printf("arraywithsubs: %v\n", mm)

	fmt.Printf("subkeyone: %v\n", mm["subkeyone"])
	fmt.Printf("subkeyone: %v\n", js.Get("test").Get("bignum").MustInt64())
}
