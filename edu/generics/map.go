package main

import "fmt"

func Map[T, U any](arr []T, transform func(T) U) []U {
	var arrans []U
	for _, x := range arr {
		arrans = append(arrans, transform(x))
	}
	return arrans
}

func main() {
	// Пример использования функции Map
	intArr := []int{1, 2, 3, 4, 5}
	doubledArr := Map(intArr, func(x int) int {
		return x * 2
	})
	// Выводим результат
	fmt.Println(doubledArr) // Ожидаемый вывод: [2 4 6 8 10]
}
