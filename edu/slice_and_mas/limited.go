package main

import (
	"errors"
	"fmt"
)

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	if n < 0 || nums == nil {
		return nil, errors.New("invalid input: n cannot be negative or nums cannot be nil")
	}
	var ans []int
	for _, elem := range nums {
		if elem < limit {
			ans = append(ans, elem)
			if len(ans) == n {
				return ans, nil
			}
		}
	}
	return ans, nil
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(UnderLimit(a, 3, 4))
}
