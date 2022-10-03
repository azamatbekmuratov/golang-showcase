package main

import "fmt"

func mapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))

	for k := range m {
		r = append(r, k)
	}

	return r
}

func reverse[T any](s []T) []T {
	l := len(s)
	r := make([]T, l)

	for i, ele := range s {
		r[l-i-1] = ele
	}
	return r
}

func print[T any](v T){
	fmt.Println(v)
}

func ForEach[T any](s []T, f func(ele T, i int, s []T)){
	for i, ele := range s {
		f(ele, i, s)
	}
}

// keys return the key of a map
// here m is generic using K and V
// V is contraint using any
// K is restrained using comparable i.e any type that supports != and == operation
func keys[K comparable, V any](m map[K]V) []K {
	// creating a slice of type K with length of map
	key := make([]K, len(m))
	i := 0
	for k, _ := range m {
		key[i] = k
		i++
	}

	return key
}

func main(){
	i := []int{4, 5, 2, 2}
	reverse(i)
}