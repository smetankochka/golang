package main

import (
	"fmt"
	"math"
)

func SqRoots() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	disc := math.Pow(b, 2) - 4.0*a*c
	if disc < 0 {
		fmt.Println(0, 0)
	} else if disc == 0 {
		fmt.Println(-1.0 * b / (2.0 * a))
	} else {
		x1 := (-1.0*b - math.Sqrt(disc)) / (2.0 * a)
		x2 := (-1.0*b + math.Sqrt(disc)) / (2.0 * a)
		if x1 < x2 {
			fmt.Println(x1, x2)
		} else {
			fmt.Println(x2, x1)
		}
	}
}
