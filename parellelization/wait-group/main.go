package main

import (
	"fmt"
	"sync"
)

func generateNumbers(total int, ch chan<- int) {

	for idx := 1; idx <= total; idx++ {
		fmt.Printf("Generating number %d\n", idx)
		ch <- idx
	}
}

func printNumbers(idx int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Printf("channel %d read value %d from channel\n", idx, num)
	}
}

func main() {
	var wg sync.WaitGroup
	numberChan := make(chan int)

	for idx := 1; idx <= 3; idx++ {
		wg.Add(1)
		go printNumbers(idx, numberChan, &wg)
	}
	generateNumbers(5, numberChan)

	close(numberChan)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}
