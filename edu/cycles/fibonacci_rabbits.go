package main

import "fmt"

func main() {
	var num, feb1, feb2 int
	fmt.Scanln(&num)
	feb1, feb2 = 1, 1
	for feb2 < num {
		feb1, feb2 = feb2, feb1+feb2
	}
	for i := 0; i < 10; i++ {
		fmt.Println(feb2)
		feb1, feb2 = feb2, feb1+feb2
	}
}
