package main

import (
	"fmt"
	"log"
)

func divide(a, b int) int {
	return a / b
}

func divideByZero() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	fmt.Println(divide(1, 0))
}

func main() {
	divideByZero()
	fmt.Println("we survived dividing by zero!")
}
