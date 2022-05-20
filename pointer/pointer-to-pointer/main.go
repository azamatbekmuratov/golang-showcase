// Go program to illustrate the
// concept of the Pointer to Pointer
package main

import "fmt"

func main() {

	// taking a variable
	// of integer type
	var v int = 100

	// taking a pointer
	// of integer type
	var pt1 *int = &v

	// taking pointer to
	// pointer to pt1
	// storing the address
	// of pt1 into pt2
	var pt2 **int = &pt1

	fmt.Println("The Value of Variable v is = ", v)

	// changing the value of v by assigning
	// the new value to the pointer pt1
	*pt1 = 200

	fmt.Println("Value stored in v after changing pt1 = ", v)

	// changing the value of v by assigning
	// the new value to the pointer pt2
	**pt2 = 300

	fmt.Println("Value stored in v after changing pt2 = ", v)

}
