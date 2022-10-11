package main

import (
	"fmt"
	"sync"
)

func f(wg *sync.WaitGroup) {
	fmt.Println("Working...")

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go f(&wg)
	wg.Wait()
	fmt.Println("Done working!")
}
