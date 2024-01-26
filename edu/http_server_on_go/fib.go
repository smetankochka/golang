package main

import (
	"fmt"
	"net/http"
)

type fibo struct {
	a, b int
}

var f fibo
var c int

func fib(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	if p == "/metrics" {
		fmt.Fprint(w, "rpc_count ", c)
		return
	}
	fmt.Fprint(w, f.a)
	tmp := f.a
	f.a = f.b
	f.b += tmp
	c++
}

func main() {
	c = 0
	f.a = 0
	f.b = 1
	http.HandleFunc("/", fib)
	http.ListenAndServe(":8080", nil)
}
