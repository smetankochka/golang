package main

import (
	"fmt"
)

func drawFlower(petalHeight int, stemHeight int) {
	for i := petalHeight; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 1; i < stemHeight; i++ {
		fmt.Println("*")
	}
	if petalHeight == 0 && stemHeight != 0 {
		fmt.Println("*")
	}
}
func main() {
	var petalHeight, stemHeight int
	fmt.Scanln(&petalHeight)
	fmt.Scanln(&stemHeight)
	drawFlower(petalHeight, stemHeight)
}
