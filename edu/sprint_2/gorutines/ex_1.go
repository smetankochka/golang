package main

import (
	"fmt"
)

func doSomething() {
	fmt.Println("hello world")
}

func main() {
	go doSomething()
}

//   !!NOTHING!!
