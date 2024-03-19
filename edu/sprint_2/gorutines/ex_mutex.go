package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	data := make(map[string]int)
	var mu sync.Mutex

	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			data["key"] = i
			mu.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			fmt.Println(data["key"])
			mu.Unlock()
		}
	}()

	time.Sleep(5 * time.Second) // Ожидание завершения работы горутин
}
