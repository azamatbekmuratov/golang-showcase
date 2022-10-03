package main

type MyStruct[T any] struct {
	inner T
}

func (m *MyStruct[T]) Get() T {
	return m.inner
}

func (m *MyStruct[T]) Set(v T) {
	m.inner = v
}