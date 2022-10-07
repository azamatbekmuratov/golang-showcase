// using atomic for primitive data types
package main

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

type WatchDog struct{ last int64 }

func (w *WatchDog) KeepAlive() {
	atomic.StoreInt64(&w.last, time.Now().Unix())
}

func (w *WatchDog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			if atomic.LoadInt64(&w.last) < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}
