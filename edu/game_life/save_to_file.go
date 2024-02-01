package main

import "os"

type World struct {
	Height int // высота сетки
	Width  int // ширина сетки
	Cells  [][]bool
}

func (w *World) SaveState(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			if w.Cells[i][j] {
				file.Write([]byte("1"))
			} else {
				file.Write([]byte("0"))
			}
		}
		if i != w.Height-1 {
			file.Write([]byte("\n"))
		}
	}
	return nil
}
