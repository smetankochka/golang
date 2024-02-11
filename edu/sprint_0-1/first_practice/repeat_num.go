package main

import "fmt"

func FindValue(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

func main() {
	fmt.Println(FindValue([]int{2, 2, 1}))
}
