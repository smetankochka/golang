package main

import "sort"

func SortSlice(slice []int) []int {
	sort.Ints(slice)
	return slice
}
