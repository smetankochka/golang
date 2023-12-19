package main

import (
	"errors"
	"fmt"
)

func GetCharacterAtPosition(str string, position int) (rune, error) {
	strRunes := []rune(str)
	if position < 0 || position >= len(strRunes) {
		return 0, errors.New("position out of range")
	}
	return strRunes[position], nil
}

func main() {
	fmt.Println(GetCharacterAtPosition("Зенит чемпион!", 20))
}
