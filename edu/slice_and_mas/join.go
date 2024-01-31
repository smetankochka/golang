package main

func Join(nums1, nums2 []int) []int {
	ans := make([]int, len(nums1)+len(nums2))
	i := 0
	for _, x := range nums1 {
		ans[i] = x
		i++
	}
	for _, x := range nums2 {
		ans[i] = x
		i++
	}
	return ans
}
