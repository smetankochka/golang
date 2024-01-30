package main

import (
	"context"
	"fmt"
	"net/http"
)

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID") // Получаем значение заголовка "X-Request-ID" из запроса
		ctx := r.Context()
		if requestID != "" {
			ctx = context.WithValue(ctx, "RequestID", requestID) // Добавляем значение в контекст запроса
		}
		next.ServeHTTP(w, r.WithContext(ctx)) // Передаем управление следующему обработчику
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	requestID := r.Context().Value("RequestID") // Получаем значение из контекста запроса
	if requestID != nil {
		fmt.Fprintf(w, "Hello! RequestID: %s", requestID)
	} else {
		fmt.Fprintf(w, "Hello! RequestID not found")
	}
}

func main() {
	http.Handle("/hello", RequestIDMiddleware(http.HandlerFunc(HelloHandler)))

	http.ListenAndServe(":8080", nil)
}
