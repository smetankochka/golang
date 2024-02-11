package main

import "fmt"

func main() {
	fmt.Println("Начало")
	panic("Что-то пошло не так") // division by zero or else
	fmt.Println("Конец")
}
