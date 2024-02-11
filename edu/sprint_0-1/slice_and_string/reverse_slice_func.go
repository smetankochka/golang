package main

func ReverseSlice(slice []int) []int {
	length := len(slice)
	new_slice := make([]int, length)
	for i := 0; i < length; i++ {
		new_slice[i] = slice[length-1-i]
	}
	return new_slice
}
