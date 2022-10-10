package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Hello, go world")

	myMap := make(map[string]int)
	var mutex = &sync.Mutex{}

	for _, c := range "abcaaaaaaaabb" {
		wg.Add(1)
		go func(letter string, hm map[string]int, wg *sync.WaitGroup) {
			defer wg.Done()
			r := rand.Intn(10)
			time.Sleep(time.Duration(r) * time.Second)
			mutex.Lock()
			hm[letter] = hm[letter] + 1
			mutex.Unlock()
		}(string(c), myMap, &wg)
	}

	wg.Wait()

	for k, v := range myMap {
		fmt.Println(k, v)
	}
}
