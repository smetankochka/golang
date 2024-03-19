package main

import (
	"sync"
	"time"
)

var (
	counter = 0
	mu      sync.Mutex
)

func incrementCounter() {
	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	go incrementCounter()
	go incrementCounter()
	go incrementCounter()
	go incrementCounter()
	time.Sleep(time.Second)
}
