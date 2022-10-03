package main

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

// Type with underlying int
type Point int

func Min[T Number](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func main() {
	// creating Point type
	x, y := Point(5), Point(2)

	fmt.Println(Min(x, y))
}