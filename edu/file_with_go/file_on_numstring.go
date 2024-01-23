package main

import (
	"bufio"
	"os"
)

func LineByNum(inputFilename string, lineNum int) string {
	file, err := os.Open(inputFilename)
	if err != nil {
		return ""
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	currentLine := 0
	for scanner.Scan() {
		if currentLine == lineNum {
			return scanner.Text()
		}
		currentLine++
	}
	return ""
}
