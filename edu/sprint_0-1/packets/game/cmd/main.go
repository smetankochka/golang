package main

import (
	"context"
	"errors"
	"github.com/smetankochka/golang/edu/sprint_0-1/packets/game/pkg/application"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	// Exit завершает программу с заданным кодом
	os.Exit(mainWithExitCode(ctx))
}

func mainWithExitCode(ctx context.Context) int {
	cfg := application.Config{
		Width:  10,
		Height: 10,
	}
	app := application.New(cfg)
	// Запускаем приложение
	if err := app.Run(ctx); err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			log.Println("Processing cancelled.")
		default:
			log.Println("Application run error", err)
		}
		// Возвращаем значение, не равное нулю, чтобы обозначить ошибку
		return 1
	}
	// Выход без ошибок
	return 0
}
