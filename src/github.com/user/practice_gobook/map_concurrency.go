package main

/*
import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter = struct {
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			counter.RLock()
			n := counter.m["key"]
			counter.RUnlock()
			fmt.Println("key:", n)
			runtime.Gosched()
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			counter.Lock()
			counter.m["key"]++
			counter.Unlock()
			runtime.Gosched()
		}
		done <- true
	}()

	<- done
	<- done
}
*/
