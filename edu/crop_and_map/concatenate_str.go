package main

import (
	"fmt"
	"strings"
)

func ConcatenateStrings(str1, str2 string) string {
	sl := []string{str1, str2}
	return strings.Join(sl, " ")
}
func main() {
	fmt.Println(ConcatenateStrings("aba", "aca"))
}
