package life_test

import (
	"github.com/smetankochka/golang/edu/sprint_0-1/packets/game/pkg/life"
	"testing"
)

func TestNewWorld(t *testing.T) {
	// Задаём размеры сетки
	height := 10
	width := 4
	// Вызываем тестируемую функцию
	world := life.NewWorld(height, width)
	// Проверяем, что в объекте указана верная высота сетки
	if world.Height != height {
		t.Errorf("expected height: %d, actual height: %d", height, world.Height)
	}
	// Проверяем, что в объекте указана верная ширина сетки
	if world.Width != width {
		t.Errorf("expected width: %d, actual width: %d", width, world.Width)
	}
	// Проверяем, что у реальной сетки — заданная высота
	if len(world.Cells) != height {
		t.Errorf("expected height: %d, actual number of rows: %d", height, len(world.Cells))
	}
	// Проверяем, что у каждого элемента — заданная длина
	for i, row := range world.Cells {
		if len(row) != width {
			t.Errorf("expected width: %d, actual row's %d len: %d", width, i, world.Width)
		}
	}
}
