package main

import (
	"bufio"
	"errors"
	"os"
)

type World struct {
	Height int // высота сетки
	Width  int // ширина сетки
	Cells  [][]bool
}

func (w *World) LoadState(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines [][]bool
	scanner := bufio.NewScanner(file)

	var width int
	var height int
	firstLine := true

	for scanner.Scan() {
		line := scanner.Text()
		if firstLine {
			width = len(line)
			firstLine = false
		} else {
			if len(line) != width {
				return errors.New("inconsistent row length in the file")
			}
		}

		var row []bool
		for _, char := range line {
			if char == '1' {
				row = append(row, true)
			} else if char == '0' {
				row = append(row, false)
			} else {
				return errors.New("invalid character in the file")
			}
		}
		lines = append(lines, row)
		height++
	}

	w.Width = width
	w.Height = height
	w.Cells = lines

	return nil
}
