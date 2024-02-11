package main

func SumOfSlice(slice []int) int {
	length := len(slice)
	sum := 0
	for i := 0; i < length; i++ {
		sum += slice[i]
	}
	return sum
}
