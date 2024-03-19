package main

import (
	"fmt"
	"time"
)

func sleeper() {
	time.Sleep(time.Second * 2)
}

func main() {
	go sleeper()
	fmt.Println("Hello, World!")
}
