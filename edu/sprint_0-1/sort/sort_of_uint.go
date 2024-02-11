package main

import (
	"fmt"
	"sort"
)

func SortNums(nums []uint) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
}

func main() {
	// Пример использования функции SortNums
	nums := []uint{5, 3, 9, 1, 7, 2, 8, 4, 6}
	fmt.Println("Исходный слайс:", nums)

	SortNums(nums)
	fmt.Println("Отсортированный слайс:", nums)
}
