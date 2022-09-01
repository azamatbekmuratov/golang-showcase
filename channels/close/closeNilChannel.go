// closing nil channels returns panic
package main

func main() {
	var c chan string
	close(c)
}
