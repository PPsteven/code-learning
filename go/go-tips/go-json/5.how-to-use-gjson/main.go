package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func main() {
	s := `{
	  "name": {"first": "Tom", "last": "Anderson"},
	  "age":37,
	  "children": ["Sara","Alex","Jack"],
	  "fav.movie": "Deer Hunter",
	  "friends": [
		{"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
		{"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
		{"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
	  ]
	}`

	paths := []string{
		"name.last",
		"age",
		"children",
		"children.#",
		"children.1",
		"child*.2",
		"c?ildren.0",
		"fav\\.movie",
		"friends.#.first",
		"friends.1.last",

		"friends.#(last==\"Murphy\").first",
		"friends.#(last==\"Murphy\")#.first",
		"friends.#(age>45)#.last",
		"friends.#(first%\"D*\").last",
		"friends.#(first!%\"D*\").last",
		"friends.#(nets.#(==\"fb\"))#.first",
	}
	for _, path := range paths {
		fmt.Printf("%s >> \t%v\n", path, gjson.Get(s, path))
	}
	//"name.last"          >> "Anderson"
	//"age"                >> 37
	//"children"           >> ["Sara","Alex","Jack"]
	//"children.#"         >> 3
	//"children.1"         >> "Alex"
	//"child*.2"           >> "Jack"
	//"c?ildren.0"         >> "Sara"
	//"fav\.movie"         >> "Deer Hunter"
	//"friends.#.first"    >> ["Dale","Roger","Jane"]
	//"friends.1.last"     >> "Craig"

	//friends.#(last=="Murphy").first    >> "Dale"
	//friends.#(last=="Murphy")#.first   >> ["Dale","Jane"]
	//friends.#(age>45)#.last            >> ["Craig","Murphy"]
	//friends.#(first%"D*").last         >> "Murphy"
	//friends.#(first!%"D*").last        >> "Craig"
	//friends.#(nets.#(=="fb"))#.first   >> ["Dale","Roger"]
}
