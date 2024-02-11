package main

func FindMaxKey(m map[int]int) int {
	maxk := -1000000000
	for key, _ := range m {
		if key > maxk {
			maxk = key
		}
	}
	return maxk
}
