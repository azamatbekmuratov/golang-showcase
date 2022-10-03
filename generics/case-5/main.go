package main

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// T as a type param now supports every int, float type
// To able to perform these operation the constrain should be only implementing types that support arthemetic operations
func Min[T Number](x, y T) T {
	if x < y {
		return x
	}
	return y
}