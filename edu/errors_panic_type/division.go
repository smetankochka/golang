package main

import (
	"errors"
	"fmt"
)

func DivideIntegers(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	ans := float64(a) / float64(b)
	return ans, nil
}

func main() {
	fmt.Println(DivideIntegers(1, 2))
}
