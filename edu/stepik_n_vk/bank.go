package main

import "fmt"
import "math"

func main() {
	var x, y int
	var p float64
	year := 0
	fmt.Scan(&x, &p, &y)
	p /= 100
	p += 1
	for x < y {
		year++
		x = int(math.Round(float64(x) * p))
	}
	fmt.Println(year)
}
