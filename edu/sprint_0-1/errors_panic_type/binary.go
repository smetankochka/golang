package main

import (
	"errors"
	"fmt"
	"strconv"
)

func IntToBinary(num int) (string, error) {
	if num < 0 {
		return "", errors.New("negative numbers are not allowed")
	}
	return strconv.FormatInt(int64(num), 2), nil
}

func main() {
	fmt.Println(IntToBinary(5))
}
