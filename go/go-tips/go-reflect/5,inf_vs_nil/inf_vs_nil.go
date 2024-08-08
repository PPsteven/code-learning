package main

import "fmt"

type messagePrinter interface {
	printMessage()
}

type thing struct {
	message string
}

func (t thing) printMessage() {
	fmt.Printf("Message: %s\n", t.message)
}

func analyzeInterface(mp interface{}) {
	fmt.Printf("Interface type: %T\n", mp)
	fmt.Printf("Interface value: %v\n", mp)
	fmt.Printf("Interface is nil: %t\n", mp == nil)
}

func main() {
	t1 := &thing{message: "hello"}
	analyzeInterface(t1)
	// Output:
	//Interface type: *main.thing
	//Interface value: &{hello}
	//Interface is nil: false

	var t2 *thing
	analyzeInterface(t2)
	// Output:
	//Interface type: *main.thing
	//Interface value: <nil>
	//Interface is nil: false

	analyzeInterface(nil)
	// Output:
	//Interface type: <nil>
	//Interface value: <nil>
	//Interface is nil: true
}
