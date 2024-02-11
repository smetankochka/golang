package main

import "fmt"

func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))

	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			merged = append(merged, left[l])
			l++
		} else {
			merged = append(merged, right[r])
			r++
		}
	}

	merged = append(merged, left[l:]...)
	merged = append(merged, right[r:]...)

	return merged
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func SortAndMerge(left, right []int) []int {
	sortedLeft := mergeSort(left)
	sortedRight := mergeSort(right)

	return merge(sortedLeft, sortedRight)
}

func main() {
	left := []int{4, 1, 3, 2}
	right := []int{8, 6, 7, 5}

	result := SortAndMerge(left, right)
	fmt.Println("Sorted and merged result:", result)
}
