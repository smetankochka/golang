package main

import "fmt"

func main() {
	var a int
	for fmt.Scan(&a); a <= 100; fmt.Scan(&a) {
		if a >= 10 {
			fmt.Println(a)
		}
	}
}
