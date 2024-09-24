package main

import (
	"fmt"
	"sync"
)

func DecrementVariable(concurrency int, decrementValue int) int {
	var mu sync.Mutex
	iter := 0
	done := false
	for {
		for i := 0; i < concurrency; i++ {
			go func() {
				mu.Lock()
				if decrementValue > 0 {
					decrementValue--
				}
				if decrementValue == 0 {
					done = true
				}
				mu.Unlock()
			}()
			if done {
				break
			}
		}
		iter++
		if done {
			break
		}
	}
	return iter
}

func main() {
	concurrency := 5
	decrementValue := 20
	iterations := DecrementVariable(concurrency, decrementValue)
	fmt.Printf("Reached zero in %d iterations.\n", iterations)
}
