package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)
	for i := 1; i < 11; i++ {
		fmt.Println(num, "*", i, "=", num*i)
	}
}
