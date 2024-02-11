package main

import (
	"errors"
	"fmt"
)

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	if n == 1 || n == 0 {
		return 1, nil
	}
	ans := 1
	for i := 2; i <= n; i++ {
		ans *= i
	}
	return ans, nil
}

func main() {
	fmt.Println(Factorial(5))
}
