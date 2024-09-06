package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	sum := a + b
	mul := a * b
	fmt.Println(sum)
	fmt.Println(mul)
}
