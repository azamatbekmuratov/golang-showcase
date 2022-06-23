// open for extension, closed for modification
package main

import "fmt"

type Cat struct {
	Name string
}

func (c Cat) Legs() int {
	return 4
}

func (c Cat) PrintLegs() {
	fmt.Printf("I have %d legs\n", c.Legs())
}

type OctoCat struct {
	Cat
}

func (oc OctoCat) Legs() int {
	return 5
}

func (oc OctoCat) PrintLegs() {
	fmt.Printf("I have %d legs\n", oc.Legs())
}

func main() {
	var octo OctoCat
	fmt.Println(octo.Legs()) // 5
	octo.PrintLegs()         // I have 5 legs
}
