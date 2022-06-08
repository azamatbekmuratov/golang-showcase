// A race condition occurs when multiple threads try to access and modify the same data (memory address)
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var sharedMapForMutex map[string]int = map[string]int{}
var mapMutex = sync.Mutex{}
var mutexReadCount int64 = 0

func runMapMutexReader(ctx context.Context, readChan chan int) {
	readCount := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("reader exiting...")
			readChan <- readCount
			return
		default:
			mapMutex.Lock()
			var _ int = sharedMapForMutex["key"]
			mapMutex.Unlock()
			readCount += 1
		}
	}
}

func runMapMutexWriter(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("writer exiting...")
			return
		default:
			mapMutex.Lock()
			sharedMapForMutex["key"] = sharedMapForMutex["key"] + 1
			mapMutex.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func startMapMutexReadWrite() {
	testContext, cancel := context.WithCancel(context.Background())

	readCh := make(chan int)
	sharedMapForMutex["key"] = 0

	numbersOfReaders := 15
	for i := 0; i < numbersOfReaders; i++ {
		go runMapMutexReader(testContext, readCh)
	}
	go runMapMutexWriter(testContext)

	time.Sleep(2 * time.Second)

	cancel()

	totalReadCount := 0
	for i := 0; i < numbersOfReaders; i++ {
		totalReadCount += <-readCh
	}

	time.Sleep(1 * time.Second)

	var counter int = sharedMapForMutex["key"]
	fmt.Printf("[MUTEX] Write Counter value: %v\n", counter)
	fmt.Printf("[MUTEX] Read Counter value: %v\n", totalReadCount)
}

func main() {
	startMapMutexReadWrite()
}
