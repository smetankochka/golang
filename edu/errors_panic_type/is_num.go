package main

import (
	"errors"
	"fmt"
	"strconv"
)

func SumTwoIntegers(a, b string) (int, error) {
	numA, errA := strconv.Atoi(a)
	if errA != nil {
		return 0, errors.New("invalid input, please provide two integers")
	}

	numB, errB := strconv.Atoi(b)
	if errB != nil {
		return 0, errors.New("invalid input, please provide two integers")
	}

	return numA + numB, nil
}

func main() {
	// Пример использования функции
	strA := "123"
	strB := "456"
	sum, err := SumTwoIntegers(strA, strB)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Sum of %s and %s is %d\n", strA, strB, sum)
	}
}
