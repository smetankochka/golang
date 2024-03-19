package main

import (
	"sync"
	"time"
)

var (
	mu    sync.Mutex
	myMap = make(map[int]int)
)

func writeToMap(key, value int) {
	mu.Lock()
	myMap[key] = value
	mu.Unlock()
}

func readFromMap(key int) {
	mu.Lock()
	c := myMap[key]
	c += 0
	mu.Unlock()
}

func main() {
	go writeToMap(0, 2)
	go writeToMap(1, 3)
	go readFromMap(1)
	go writeToMap(3, 4)
	go writeToMap(0, 3)
	go readFromMap(0)
	time.Sleep(time.Second)
}
