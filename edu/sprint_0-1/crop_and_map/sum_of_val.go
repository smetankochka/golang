package main

func SumOfValuesInMap(m map[int]int) int {
	sum := 0
	for _, val := range m {
		sum += val
	}
	return sum
}
