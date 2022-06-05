package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	buf := make([]byte, 512)
	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			fmt.Printf("buf[:n], %q\n", buf[:n])
		}
	}
}
