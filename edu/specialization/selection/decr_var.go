package main

import (
	"fmt"
	"sync"
)

func DecrementVariable(concurrency int, decrementValue *int) int {
	var mu sync.Mutex
	var wg sync.WaitGroup

	iterations := 0
	done := false

	for !done {
		wg.Add(concurrency)
		iterations++
		for i := 0; i < concurrency; i++ {
			go func() {
				defer wg.Done()
				mu.Lock()
				defer mu.Unlock()
				if *decrementValue > 0 {
					(*decrementValue)--
				}
				if *decrementValue <= 0 {
					done = true
				}
			}()
		}

		wg.Wait()
	}

	fmt.Print(iterations)
	return iterations
}

func main() {
	initialDecrementValue := 10
	iterations := DecrementVariable(3, &initialDecrementValue)
	fmt.Println("\nTotal iterations:", iterations)
}
