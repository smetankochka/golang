package main

import "fmt"

func Flatten2DArray(arr [][]int) []int {
	res := make([]int, 0)
	for _, x := range arr {
		for _, y := range x {
			res = append(res, y)
		}
	}
	return res
}

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 0, 0},
	}

	res := Flatten2DArray(arr)
	fmt.Print(res)
}
