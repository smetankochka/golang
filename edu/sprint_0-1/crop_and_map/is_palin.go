package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(input string) bool {
	input = strings.ToLower(input)
	var filteredString1 string
	for _, r := range input {
		if r != ' ' {
			filteredString1 += string(r)
		}
	}
	filteredString := []rune(filteredString1)
	l := len(filteredString)
	for i := 0; i < l/2; i++ {
		if filteredString[i] != filteredString[l-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome("А роза упала на лапу Азора")) // вернет true
	fmt.Println(IsPalindrome("А роза упала на лапу Азор"))  // вернет false
}
