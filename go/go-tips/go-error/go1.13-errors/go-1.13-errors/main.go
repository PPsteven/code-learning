package main

import (
	"fmt"
	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("not found")

type NotFoundError struct {
	Name string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("name: %s not found", e.Name)
}

func FindByName(name string) error {
	//return errors.Wrapf(ErrNotFound, "name: %s not found", name)
	return fmt.Errorf("name: %s not found: %w", name, ErrNotFound)
	// %w in fmt.Errorf is to wrap err to a new error
}

func FindByName2(name string) error {
	return &NotFoundError{Name: name}
}

func main() {
	err := FindByName("Tom")
	if errors.Is(err, ErrNotFound) {
		fmt.Println(err)
		// name: Tom not found: not found
	}

	err = FindByName2("Tom")
	var e *NotFoundError
	if errors.As(err, &e) {
		fmt.Println(e)
		// name: Tom not found
	}
}
