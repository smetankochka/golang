package main

import "fmt"

func main() {
	var first, second int
	fmt.Scanln(&first, &second)
	if first > 0 && second > 0 {
		fmt.Println("Оба числа положительные")
	} else if first < 0 && second < 0 {
		fmt.Println("Оба числа отрицательные")
	} else if first == 0 || second == 0 {
		fmt.Println("Одно из чисел равно нулю")
	} else {
		fmt.Println("Одно число положительное, а другое отрицательное")
	}
}
