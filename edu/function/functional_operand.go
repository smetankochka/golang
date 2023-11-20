package main

import (
	"fmt"
	"math"
)

func QuadraticEquationRoots(a, b, c float64) (float64, float64) {
	disc := math.Pow(b, 2) - 4.0*a*c
	if disc < 0 {
		return math.NaN(), math.NaN()
	} else if disc == 0 {
		return -1.0 * b / (2.0 * a), -1.0 * b / (2.0 * a)
	} else {
		x1 := (-1.0*b - math.Sqrt(disc)) / (2.0 * a)
		x2 := (-1.0*b + math.Sqrt(disc)) / (2.0 * a)
		if x1 < x2 {
			return x1, x2
		} else {
			return x2, x1
		}
	}
}

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) float64 {
	if b == 0 {
		return math.NaN()
	}
	return a / b
}

func PrintNumbersAscending(n int) {
	for i := 1; i <= n; i++ {
		fmt.Print(i, " ")
	}
}

func PrintNumbersDescending(n int) {
	for i := n; i > 0; i-- {
		fmt.Print(i, " ")
	}
}
