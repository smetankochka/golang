package main

import (
	"fmt"
	"sync"
	"time"
)

type VisitCounter struct {
	mu       sync.Mutex
	visits   map[string]int
	lastSeen map[string]time.Time
	ttl      time.Duration
}

func NewVisitCounter(ttl time.Duration) *VisitCounter {
	vc := &VisitCounter{
		visits:   make(map[string]int),
		lastSeen: make(map[string]time.Time),
		ttl:      ttl,
	}

	go vc.cleanupExpiredVisits()

	return vc
}

func (vc *VisitCounter) RecordVisit(user string) {
	vc.mu.Lock()
	defer vc.mu.Unlock()

	vc.visits[user]++
	vc.lastSeen[user] = time.Now()
}

func (vc *VisitCounter) GetVisitCount(user string) int {
	vc.mu.Lock()
	defer vc.mu.Unlock()
	if vc.isExpired(user) {
		delete(vc.visits, user)
		delete(vc.lastSeen, user)
		return 0
	}

	return vc.visits[user]
}

func (vc *VisitCounter) SetTTL(ttl time.Duration) {
	vc.mu.Lock()
	defer vc.mu.Unlock()

	vc.ttl = ttl
}

func (vc *VisitCounter) cleanupExpiredVisits() {
	for {
		time.Sleep(1 * time.Second)

		vc.mu.Lock()
		now := time.Now()
		for user, lastSeen := range vc.lastSeen {
			if now.Sub(lastSeen) > vc.ttl {
				delete(vc.visits, user)
				delete(vc.lastSeen, user)
			}
		}
		vc.mu.Unlock()
	}
}

func (vc *VisitCounter) isExpired(user string) bool {
	lastSeen, exists := vc.lastSeen[user]
	if !exists {
		return true
	}
	return time.Since(lastSeen) > vc.ttl
}

func main() {
	ttl := 10 * time.Second
	vc := NewVisitCounter(ttl)

	vc.RecordVisit("Alice")
	vc.RecordVisit("Alice")

	fmt.Printf("Alice visit count: %d\n", vc.GetVisitCount("Alice")) // Ожидается: 2
	time.Sleep(2 * time.Second)

	time.Sleep(10 * time.Second) // Ждем, чтобы TTL истек

	fmt.Printf("Alice visit count after TTL: %d\n", vc.GetVisitCount("Alice")) // Ожидается: 0
}
