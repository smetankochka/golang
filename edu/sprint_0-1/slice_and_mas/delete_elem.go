package main

func Clean(nums []int, x int) []int {
	i := 0
	for _, num := range nums {
		if num != x {
			nums[i] = num
			i++
		}
	}
	return nums[:i]
}
