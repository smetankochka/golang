package main

import "fmt"

func Rotate(data []int, pos int) []int {
	l := len(data)
	if l == 0 {
		return data
	}
	pos = (pos%l + l) % l
	result := make([]int, l)
	copy(result, append(data[l-pos:], data[:l-pos]...))
	return result
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	pos := -3
	result := Rotate(data, pos)
	fmt.Println(result) // Output: [5 6 7 1 2 3 4]
}
