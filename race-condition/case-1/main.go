package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu sync.Mutex
)

func main() {
	numbers := []int{}
	go appendOne(&numbers)
	go appendTwo(&numbers)
	time.Sleep(30 * time.Millisecond)
}

func appendOne(numbers *[]int) {
	mu.Lock()
	defer mu.Unlock()
	*numbers = append(*numbers, 12)
}

func appendTwo(numbers *[]int) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Printf("numbers %v\n", numbers)
	*numbers = append(*numbers, []int{1, 2}...)
	fmt.Printf("numbers %v\n", numbers)
}
