// Joining strings in slice
package main

import (
	"fmt"
	"strings"
)

func main() {
	mySlice := []string{"Welcome", "to", "golang"}

	result := strings.Join(mySlice, "-")
	fmt.Println(result)
}
