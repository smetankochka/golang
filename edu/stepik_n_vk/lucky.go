package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var s1, s2 int
	for i := 0; i < 3; i++ {
		s1 += a % 10
		a /= 10
	}
	for i := 0; i < 3; i++ {
		s2 += a % 10
		a /= 10
	}
	if s1 == s2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
