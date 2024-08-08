package main

import (
	"fmt"
	"reflect"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func main() {
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
	}

	// 先传指针，后取实例，是为了能够修改
	valueOfStrangelove := reflect.ValueOf(&strangelove)
	valueOfStrangelove = valueOfStrangelove.Elem()

	fmt.Printf("before: %v\n", strangelove)
	valueOfStrangelove.FieldByName("Title").SetString("Harry Potter")
	valueOfStrangelove.Field(2).Set(reflect.ValueOf(2010))
	fmt.Printf("after: %v\n", strangelove)
	// before: {Dr. Strangelove How I Learned to Stop Worrying and Love the Bomb 1964 false map[] [] <nil>}
	// after: {Harry Potter How I Learned to Stop Worrying and Love the Bomb 2010 false map[] [] <nil>}

}
