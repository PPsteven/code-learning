package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type NotFoundError struct {
	Name string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("name: %s not found", e.Name)
}

func FindByName(name string) error {
	if len(name) > 0 {
		return fmt.Errorf("name: %s not found", name)
	}
	return ErrNotFound
}

func FindByName2(name string) error {
	return &NotFoundError{Name: name}
}

func main() {
	err := FindByName("")
	if err == ErrNotFound {
		fmt.Println(err)
		// not found
	}

	err = FindByName("Jack")
	// easy to read by human but not program
	if err != nil && err.Error() == fmt.Sprintf("name: %s not found", "Jack") {
		fmt.Println(err)
		// name: Jack not found
	}

	err = FindByName2("Jack")
	if err != nil {
		// easy to read by program
		if e, ok := err.(*NotFoundError); ok {
			fmt.Println(e)
			// name: Jack not found
		}
	}
}

