package main

func Mix(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums)/2; i++ {
		ans[i*2] = nums[i]
		ans[2*2+1] = nums[i+len(nums)/2]
	}
	return ans
}
