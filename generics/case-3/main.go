package main

// Generic struct with two generic types
type Enteries[K, V any] struct {
	Key K
	Value V
}

func enteries[K comparable, V any](m map[K]V) []*Enteries[K, V] {
	// define a slice with Enteries type passing, K, V type parameters
	e := make([]*Enteries[K,V], len(m))
	i := 0

	for k, v := range m {
		// creating value using new keyword
		newEntery := new(Enteries[K, V])
		newEntery.Key = k
		newEntery.Value = v
		e[i] = newEntery
		i++
	}
	return e
}