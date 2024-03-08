package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func SumUp(filepath, colname string) (int, error) {
	// Открываем файл
	file, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Читаем CSV
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return 0, err
	}
	if len(records) == 0 {
		return 0, fmt.Errorf("в указанной колонке нет значений")
	}
	// Находим индекс колонки по имени
	var colIndex int
	colIndex = -1
	for i, name := range records[0] {
		if name == colname {
			colIndex = i
			break
		}
	}
	if colIndex == -1 {
		return 0, fmt.Errorf("в указанной колонке нет значений")
	}
	// Проверяем, что колонка не пуста
	if len(records) < 2 {
		return 0, fmt.Errorf("в указанной колонке нет значений")
	}

	// Считаем сумму значений в указанной колонке
	sum := 0
	for _, row := range records[1:] {
		val, err := strconv.Atoi(row[colIndex])
		if err != nil {
			return 0, err
		}
		sum += val
	}

	return sum, nil
}

func main() {
	sum, err := SumUp("data.csv", "colname")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Сумма значений из колонки:", sum)
	}
}
