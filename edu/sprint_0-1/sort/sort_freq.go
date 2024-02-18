package main

import (
	"fmt"
	"sort"
)

type CharFreq struct {
	Char  byte
	Freq  int
	Index int
}

func SortByFreq(str string) string {
	freqMap := make(map[byte]int)
	var charFreqs []CharFreq

	for i := 0; i < len(str); i++ {
		freqMap[str[i]]++
	}

	for char, freq := range freqMap {
		charFreqs = append(charFreqs, CharFreq{Char: char, Freq: freq, Index: int(char)})
	}

	sort.Slice(charFreqs, func(i, j int) bool {
		if charFreqs[i].Freq == charFreqs[j].Freq {
			return charFreqs[i].Index < charFreqs[j].Index
		}
		return charFreqs[i].Freq < charFreqs[j].Freq
	})

	result := ""
	for _, cf := range charFreqs {
		for i := 0; i < cf.Freq; i++ {
			result += string(cf.Char)
		}
	}

	return result
}

func main() {
	input := "abbbzzzat"
	fmt.Println("Input:", input)
	fmt.Println("Output:", SortByFreq(input))
}
