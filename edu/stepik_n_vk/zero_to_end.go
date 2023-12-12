package main

import "fmt"

func main() {
	var a, b int
	count := 0
	fmt.Scan(&a)
	sl := make([]int, 0)
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		if b != 0 {
			sl = append(sl, b)
		} else {
			count++
		}
	}
	for _, val := range sl {
		fmt.Print(val, " ")
	}
	for i := 0; i < count; i++ {
		fmt.Print(0, " ")
	}
}
