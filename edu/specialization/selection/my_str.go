package main

import (
	"fmt"
	"sync"
)

type VisitCounter struct {
	mu     sync.RWMutex
	visits map[string]int
}

func NewVisitCounter() *VisitCounter {
	return &VisitCounter{
		visits: make(map[string]int),
	}
}

func (vc *VisitCounter) RecordVisit(user string) {
	vc.mu.Lock()
	defer vc.mu.Unlock()
	vc.visits[user]++
}

func (vc *VisitCounter) GetVisitCount(user string) (int, bool) {
	vc.mu.RLock()
	defer vc.mu.RUnlock()
	count, exists := vc.visits[user]
	return count, exists
}

func (vc *VisitCounter) DeleteUser(user string) {
	vc.mu.Lock()
	defer vc.mu.Unlock()
	delete(vc.visits, user)
}

func main() {
	vc := NewVisitCounter()

	// Примеры использования
	vc.RecordVisit("Alice")
	vc.RecordVisit("Alice")

	count, exists := vc.GetVisitCount("Alice")
	if exists {
		fmt.Printf("Алиса зашла на сайт %d раз.\n", count)
	} else {
		fmt.Println("Алиса не найдена в системе.")
	}

	vc.DeleteUser("Alice")
	count, exists = vc.GetVisitCount("Alice")
	if exists {
		fmt.Printf("Алиса зашла на сайт %d раз.\n", count)
	} else {
		fmt.Println("Алиса не найдена в системе.")
	}
}
