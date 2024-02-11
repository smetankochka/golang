package main

func IntersectionOfSlices(slice1, slice2 []int) []int {
	l1 := len(slice1)
	l2 := len(slice2)
	new_slice := make([]int, 0)
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if slice1[i] == slice2[j] {
				new_slice = append(new_slice, slice1[i])
				break
			}
		}
	}
	return new_slice
}
