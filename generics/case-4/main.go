package main

// Stringer is a constraint
type Stringer interface {
	String() string
}

// Here T has to implement Stringer, T can only perform operations defined by Stringer
func stringer[T Stringer](s T) string {
	return s.String()
}