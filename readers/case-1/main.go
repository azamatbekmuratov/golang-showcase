package main

import (
	"io"
	"strings"
	"fmt"
)

func main() {
	r := strings.NewReader("It is a reader")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n=%v err=%err b=%b\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
