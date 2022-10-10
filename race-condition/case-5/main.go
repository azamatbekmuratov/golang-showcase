package main

import "sync"

type SafeNumber struct {
	val int
	m   sync.Mutex
}

func (i *SafeNumber) Get() int {
	i.m.Lock()
	defer i.m.Unlock()
	return i.val
}

func (i *SafeNumber) Set(val int) {
	i.m.Lock()
	defer i.m.Unlock()
	i.val = val
}

func getNumber() int {
	i := &SafeNumber{}

	go func() {
		i.Set(5)
	}()

	return i.Get()
}
