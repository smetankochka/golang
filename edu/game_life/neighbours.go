package main

type World struct {
	Height int // высота сетки
	Width  int // ширина сетки
	Cells  [][]bool
}

func (w *World) Neighbors(x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue // пропускаем текущую клетку
			}
			// Проверяем "живая" ли соседняя клетка (учитываем граничные условия, чтобы не выйти за пределы сетки)
			if y+i >= 0 && y+i < w.Height && x+j >= 0 && x+j < w.Width && w.Cells[y+i][x+j] {
				count++
			}
		}
	}
	return count
}
