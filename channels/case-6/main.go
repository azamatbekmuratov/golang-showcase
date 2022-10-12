package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)
	ch <- "geeksforgeeks"
	ch <- "geeksforgeeks world"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
