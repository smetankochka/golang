package main

import "fmt"

func main() {
	var a, b int
	sum := 0
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		if b%8 == 0 && b > 9 && b < 100 {
			sum += b
		}
	}
	fmt.Print(sum)
}
