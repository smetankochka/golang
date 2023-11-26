package main

import (
	"fmt"
)

func MoveZeroes(nums []int) []int {
	result := []int{}
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			result = append(result, nums[i])
		} else {
			count++
		}
	}
	for i := 0; i < count; i++ {
		result = append(result, 0)
	}
	return result
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	mas := MoveZeroes(nums)
	for i := 0; i < len(mas); i++ {
		fmt.Println(mas[i])
	}
}
