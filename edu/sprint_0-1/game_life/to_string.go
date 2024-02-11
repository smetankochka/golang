package main

import "strings"

type World struct {
	Height int // высота сетки
	Width  int // ширина сетки
	Cells  [][]bool
}

func (w World) String() string {
	var builder strings.Builder
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			if w.Cells[i][j] {
				// Живая клетка будет представлена зеленым квадратом
				builder.WriteString("\xF0\x9F\x9F\xA9") // зеленый квадрат
			} else {
				// Мертвая клетка будет представлена коричневым квадратом
				builder.WriteString("\xF0\x9F\x9F\xAB")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}
