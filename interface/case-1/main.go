package main

import "fmt"

type Squarer interface {
	GetSquare() int
}

type Foursquare struct {
	a int
}

func (obj Foursquare) GetSquare() int {
	return obj.a * obj.a
}

type Triangle struct {
	a int
	b int
	c int
}

func (obj Triangle) GetSquare() int {
	return obj.a * obj.b / 2
}

func sumSquare(s []Squarer) int {
	square := 0

	for i := range s {
		square += s[i].GetSquare()
	}

	return square
}

func main() {
	figues := []Squarer{Foursquare{a: 3}, Foursquare{a: 2}, Triangle{a: 2, b: 3}}

	fmt.Println(sumSquare(figues))
}
