package main

import "fmt"

func MaxExpressionValue(nums []int) int {
	maxVal := -1 // Используем начальное невозможно низкое значение
	for p := 0; p < len(nums); p++ {
		for q := p + 1; q < len(nums); q++ {
			for r := q + 1; r < len(nums); r++ {
				for s := r + 1; s < len(nums); s++ {
					if nums[s]-nums[r]+nums[q]-nums[p] > maxVal {
						maxVal = nums[s] - nums[r] + nums[q] - nums[p]
					}
				}
			}
		}
	}
	return maxVal
}

func main() {
	nums := []int{4, 5, 10, 50, 25}
	fmt.Println(MaxExpressionValue(nums)) // Это должно вывести 46
}
