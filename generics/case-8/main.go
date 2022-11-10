package main

type customConstraint interface {
	~int | ~string 
}

func getKeys(K customConstraint, V any)(m map[K]V) []K {
	var keys []K
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}

func main() {

}