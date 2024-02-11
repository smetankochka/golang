package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	sum := a + b
	fmt.Println("Сумма чисел:", sum)
}
