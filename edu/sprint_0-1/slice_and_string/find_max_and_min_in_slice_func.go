package main

func FindMinMaxInSlice(slice []int) (int, int) {
	length := len(slice)
	if length == 0 {
		return 0, 0
	}
	minimum := slice[0]
	maximum := slice[0]
	for i := 1; i < length; i++ {
		if slice[i] < minimum {
			minimum = slice[i]
		}
		if slice[i] > maximum {
			maximum = slice[i]
		}
	}
	return minimum, maximum
}
