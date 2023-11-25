package main

import (
	"fmt"
)

func FindMissingValues(nums []int) []int {
	flag := make([]bool, len(nums))
	result := []int{}
	for i := 0; i < len(flag); i++ {
		flag[i] = false
	}
	for i := 0; i < len(flag); i++ {
		flag[nums[i]-1] = true
	}
	for i := 0; i < len(flag); i++ {
		if flag[i] == false {
			result = append(result, i+1)
		}
	}
	return result
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	mas := FindMissingValues(nums)
	for i := 0; i < len(mas); i++ {
		fmt.Println(mas[i])
	}
}
