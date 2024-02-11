package main

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Логируем информацию о запросе
		log.Printf("Запрос: %s %s", r.Method, r.URL.Path)

		// Передаем управление следующему обработчику
		next.ServeHTTP(w, r)

		// Вычисляем время выполнения запроса
		duration := time.Since(start)
		log.Printf("Время выполнения запроса: %s", duration)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет, мир!"))
}

func main() {
	mux := http.NewServeMux()

	// Создаем обработчик для маршрута "/"
	hello := http.HandlerFunc(helloHandler)

	// Применяем logging middleware к обработчику "/"
	mux.Handle("/", loggingMiddleware(hello))

	// Запускаем сервер на порту 8080
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
