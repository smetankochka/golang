package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Scan(&num2, &num1)
	digitsMap := make(map[int]bool)

	// Преобразуем num1 в карту
	for num1 > 0 {
		digit := num1 % 10
		digitsMap[digit] = true
		num1 /= 10
	}

	// Проверяем цифры num2 на наличие в карте
	commonDigits := []int{}
	for num2 > 0 {
		digit := num2 % 10
		if digitsMap[digit] {
			commonDigits = append(commonDigits, digit)
		}
		num2 /= 10
	}

	// Выводим общие цифры
	for i := len(commonDigits) - 1; i >= 0; i-- {
		fmt.Printf("%d ", commonDigits[i])
	}
}
