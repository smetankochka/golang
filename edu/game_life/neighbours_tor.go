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
				continue // Пропускаем текущую клетку
			}

			// В этой отдельной функции мы обрабатываем тороидальные границы
			// Это позволит нам перепрыгивать с одного края сетки на другой, создавая эффект тора
			// Обработка тороидальных границ осуществляется с помощью взятия остатка от деления на ширину и высоту сетки
			nx := (x + i + w.Width) % w.Width
			ny := (y + j + w.Height) % w.Height

			if w.Cells[ny][nx] {
				count++
			}
		}
	}
	return count
}
