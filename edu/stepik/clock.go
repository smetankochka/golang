package main

import "fmt"

func main() {
	var d int
	fmt.Scan(&d)
	h := d / 30
	m := (d % 30) * 2
	fmt.Printf("It is %d hours %d minutes.\n", h, m)
}
