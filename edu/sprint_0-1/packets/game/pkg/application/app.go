package application

import (
	"context"
	"fmt"
	"github.com/smetankochka/golang/edu/sprint_0-1/packets/game/pkg/life"
	"time"
)

type Config struct {
	Width  int
	Height int
}

type Application struct {
	Cfg Config
}

func New(config Config) *Application {
	return &Application{
		Cfg: config,
	}
}

func (a *Application) Run(ctx context.Context) error {
	// Объект для хранения текущего состояния сетки
	currentWorld := life.NewWorld(a.Cfg.Height, a.Cfg.Width)
	// Объект для хранения очередного состояния сетки
	nextWorld := life.NewWorld(a.Cfg.Height, a.Cfg.Width)
	// Заполняем сетку на 30%
	currentWorld.RandInit(30)
	for {
		// Здесь мы можем записывать текущее состояние  — например, в очередь сообщений. Для нашего примера просто выводим на экран
		fmt.Println(currentWorld)
		life.NextState(currentWorld, nextWorld)
		currentWorld = nextWorld
		// Проверяем контекст
		select {
		case <-ctx.Done():
			return ctx.Err() // Возвращаем причину завершения
		default: // По умолчанию делаем паузу
			time.Sleep(100 * time.Millisecond)
			break
		}
		// Очищаем экран
		fmt.Print("\033[H\033[2J")
	}
}
