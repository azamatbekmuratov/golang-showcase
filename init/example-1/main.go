package main

import (
	"math/rand"
	"fmt"
)

var creatures = []string{"shark", "jellyfish", "squid", "octopus", "dolphin"}

func Random() string {
	i := rand.Intn(len(creatures))
	return creatures[i]
}

func main() {
	fmt.Println(creatures)
}
