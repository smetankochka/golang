package main

import "fmt"

func main() {
	var meters float32 = 0.0
	fmt.Scan(&meters)
	const exchangeRate = 1852
	seaMiles := meters / exchangeRate
	fmt.Println(seaMiles)
}
