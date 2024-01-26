package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func isEnglish(s string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z]*$", s)
	return match
}

func greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "hello stranger")
	} else if isEnglish(name) {
		fmt.Fprintf(w, "hello %s", name)
	} else {
		fmt.Fprint(w, "hello dirty hacker")
	}
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
