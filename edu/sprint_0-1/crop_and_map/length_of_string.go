package main

import (
	"fmt"
)

func StringLength(input string) int {
	letters := []rune(input)
	return len(letters)
}
func main() {
	fmt.Println(StringLength("abacsa"))
}
