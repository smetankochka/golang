package main

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s %s {%s} %s\n", time.Now().Format("2006/01/02 15:04:05"), r.Method, r.URL.Path, r.RemoteAddr, r.Header, time.Since(start))
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func main() {
	http.Handle("/", loggingMiddleware(http.HandlerFunc(handler)))
	log.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server error: %s", err)
	}
}
