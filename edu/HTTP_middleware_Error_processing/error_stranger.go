package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

func Sanitize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		match, _ := regexp.MatchString("^[a-zA-Z]*$", name)
		if !match {
			http.Error(w, "Invalid name", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func SetDefaultName(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "stranger"
			r.URL.RawQuery = "name=" + name
		}
		next.ServeHTTP(w, r)
	})
}

func greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", greet)
	gr := http.HandlerFunc(greet)
	finalHandler := SetDefaultName(Sanitize(gr))
	http.Handle("/", finalHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
