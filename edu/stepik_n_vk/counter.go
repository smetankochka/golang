package main

import (
	"fmt"
)

func main() {
	var a, b int
	count := 0
	fmt.Scan(&a)
	new_slice := make([]int, 0)
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		new_slice = append(new_slice, b)
	}
	fmt.Scan(&b)
	for _, val := range new_slice {
		if val != b {
			count++
		}
	}
	fmt.Print(count)
}
