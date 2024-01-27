package main

import (
	"fmt"
	"log"
	"net/http"
)

type fibo struct {
	a, b int
}

var f fibo
var requestCount int

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, f.a)
	tmp := f.a
	f.a = f.b
	f.b += tmp
	requestCount++
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "rpc_duration_milliseconds_count ", requestCount)
}

func Metrics(metr http.Handler, fib http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		if p == "/metrics" {
			metr.ServeHTTP(w, r)
		} else {
			fib.ServeHTTP(w, r)
		}
	})
}

func main() {
	requestCount = 0
	f.a = 0
	f.b = 1
	mux := http.NewServeMux()
	metr := http.HandlerFunc(MetricsHandler)
	fib := http.HandlerFunc(FibonacciHandler)
	mux.Handle("/", Metrics(metr, fib))
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
