package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Tick(1 * time.Second)
	for t := range ticker {
		fmt.Println("Tick at", t)
	}
}
