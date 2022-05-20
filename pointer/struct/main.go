// Golang program to illustrate the
// concept of the Pointer to struct
package main

import "fmt"

type Employee struct {
	name  string
	empid int
}

func main() {

	// creating the instance of the
	// Employee struct type
	emp := Employee{"ABC", 19078}

	// Here, it is the pointer to the struct
	pts := &emp

	fmt.Println(pts)

	// accessing the struct fields(liem employee's name)
	// using a pointer but here we are not using
	// dereferencing explicitly
	fmt.Println(pts.name)

	// same as above by explicitly using
	// dereferencing concept
	// means the result will be the same
	fmt.Println((*pts).name)

}
