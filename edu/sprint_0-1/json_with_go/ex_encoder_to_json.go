package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}

func main() {
	// Создаём слайс структур Student с данными о школьниках
	students := []Student{
		{Name: "Alice", Age: 12, Grade: 7},
		{Name: "Bob", Age: 13, Grade: 8},
		{Name: "Charlie", Age: 14, Grade: 9},
	}

	// Создаём буфер для записи JSON-данных
	var buf bytes.Buffer

	// Создаём Encoder для записи JSON-данных в буфер
	encoder := json.NewEncoder(&buf)

	// Записываем JSON-данные в буфер с помощью метода Encode() Encoder
	err := encoder.Encode(students)
	if err != nil {
		fmt.Println("Ошибка при записи JSON-данных:", err)
		return
	}

	// Выводим результат на экран
	fmt.Println("JSON-данные о студентах:")
	fmt.Println(buf.String())
}
