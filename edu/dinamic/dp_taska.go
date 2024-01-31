package main

func MaxExpressionValue(nums []int) int {
	n := len(nums)
	maxVal := make([][]int, n)
	for i := range maxVal {
		maxVal[i] = make([]int, n)
	}

	maxDiff := make([][]int, n)
	for i := range maxDiff {
		maxDiff[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			maxDiff[i][j] = nums[j] - nums[i]
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			maxVal[i][j] = maxDiff[i][j]
			if i > 0 && j < n-1 {
				maxVal[i][j] += maxVal[i-1][j+1]
			}
		}
	}

	result := maxVal[0][1]
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if maxVal[i][j] > result {
				result = maxVal[i][j]
			}
		}
	}

	return result
}
