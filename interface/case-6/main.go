package main

import (
	"fmt"
)

type error interface {
	Error() string
}

type MyError struct {
	description string
}

func (err MyError) Error() string {
	return fmt.Sprintf("error: %s", err.description)
}

func f() error {
	return MyError{"foo"}
}

func main() {
	fmt.Printf("%v\n", f()) // error: foo
}
