// source: gopl.io/ch12/methods

package main

import (
	"fmt"
	"reflect"
	"strings"
)

type example struct {
}

func (e *example) A(_ string) bool {
	return true
}

func (e *example) B(_ int, _ map[string][]string) {
}


func PrintSignature(x interface{}) {
	v := reflect.ValueOf(x)
	t := reflect.TypeOf(x)

	fmt.Printf("type %s\n", t.Elem())

	for i := 0; i < v.NumMethod(); i++ {
		methodTyp := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name,
			strings.TrimPrefix(methodTyp.String(), "func"))
	}
}

func main() {
	e := &example{}
	PrintSignature(e)
}

// Output:
// type main.example
// func (*main.example) A(string) bool
// func (*main.example) B(int, map[string][]string)

