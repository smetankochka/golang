package main

import "fmt"

func main() {
	var a, maximum, count int
	count = 0
	maximum = -1
	for fmt.Scan(&a); a != 0; fmt.Scan(&a) {
		if a > maximum {
			maximum = a
			count = 1
		} else if a == maximum {
			count++
		}
	}
	fmt.Print(count)
}
