package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/tidwall/gjson"
)

// 1. decode to struct when structure is fixed
// 2. get field value when structure is unfiexd by map
// 3. package: go-simplejson
// 4. package: gjson

type BodyInfo struct {
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
}

type Bird struct {
	Body BodyInfo `json:"body"`
	Name string   `json:"name"`
}

// stuctured
func readHeightFromStructuredData(birdString string) float64 {
	var birds []Bird
	_ = json.Unmarshal([]byte(birdString), &birds)
	return birds[0].Body.Height
}

// unstructured data like maps
func readHeightFromUnstructuredData(birdString string) float64 {
	var people []interface{}
	_ = json.Unmarshal([]byte(birdString), &people)
	if len(people) <= 0 {
		return 0
	}

	person := people[0] // interface
	for k, v := range person.(map[string]interface{}) {
		if k == "body" {
			for kk, vv := range v.(map[string]interface{}) {
				if kk == "height" {
					return vv.(float64)
				}
			}
		}
	}
	return 0
}

func readHeightByGoSimplejson(birdString string) float64 {
	js, _ := simplejson.NewJson([]byte(birdString))
	height, _ := js.GetIndex(0).GetPath("body", "height").Float64()
	return height
}

// use gjson "github.com/tidwall/gjson"
func readHeightByGjson(birdString string) float64 {
	return gjson.Get(birdString, "0").Get("body.height").Float()
}


func main() {
	// 1. decode to struct when structure is fixed
	birdString := `[{"body": {"height": 163.3,"weight": 32},"name": "Lily"}]`
	height := readHeightFromStructuredData(birdString)
	fmt.Printf("height: %0.1f\n", height)

	// 2. get field value when structure is unfiexd by map
	height = readHeightFromUnstructuredData(birdString)
	fmt.Printf("height: %0.1f\n", height)

	// 3. package: go-simplejson
	height = readHeightByGoSimplejson(birdString)
	fmt.Printf("height: %0.1f\n", height)

	// 4. package: gojson
	height = readHeightByGjson(birdString)
	fmt.Printf("height: %0.1f\n", height)
}
