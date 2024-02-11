package main

import "fmt"

func Convert2D(nums []int, m, n int) [][]int {
	result := make([][]int, m)
	idx := 0
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = nums[idx]
			idx++
			if idx == len(nums) {
				return result
			}
		}
	}
	return result
}

func main() {
	fmt.Println(Convert2D([]int{1, 2, 3, 4, 5, 6, 7, 8}, 4, 2))
}
