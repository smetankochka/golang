package main

import (
	"fmt"
	"time"
)

func doSomething() {
	fmt.Println("Hello, World!")
}

func main() {
	go doSomething()
	time.Sleep(1 * time.Second)
}
