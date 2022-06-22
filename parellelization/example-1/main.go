package main

import (
	"fmt"
	"runtime"
)

type Vector []float64

const numCPU = runtime.NumCPU() //number of CPU cores

func (v Vector) doAll(u Vector) {
	c := make(chan int, numCPU)
	for i := 0; i < len(numCPU); i++ {

	}
	//Drain the channel
	for i := 0; i < len(numCPU); i++ {
		<-c //wait for one task for complete
	}
	//All done
}

func main() {
	fmt.Println("Parallelization")
}
