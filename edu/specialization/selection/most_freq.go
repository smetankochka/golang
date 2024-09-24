package main

import (
	"fmt"
	"strings"
)

func MostFrequentConsonant(s string) string {
	s = strings.ToLower(s)
	counts := make(map[rune]int)
	consonants := "qwrtpsdfghjklzxcvbnm"
	for _, char := range s {
		if strings.ContainsRune(consonants, char) {
			counts[char]++
		}
	}
	var maxCount int
	var result rune
	for char, count := range counts {
		if count > maxCount || (count == maxCount && char < result) {
			maxCount = count
			result = char
		}
	}
	if maxCount == 0 {
		return ""
	}
	return string(result)
}

func main() {
	s := "aaajjjjjass"
	res := MostFrequentConsonant(s)
	fmt.Println(res) // Результат должен быть: "s"
}
