package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonStr := `{"name": "John", "age": 30}`

	// Создаём буфер с JSON-данными
	reader := strings.NewReader(jsonStr)

	// Создаём Decoder для чтения JSON из буфера
	decoder := json.NewDecoder(reader)

	// Создаём переменную для хранения декодированных данных
	var person Person

	// Читаем JSON из буфера и записываем в переменную person
	err := decoder.Decode(&person)
	if err != nil {
		fmt.Println("Ошибка чтения JSON:", err)
		return
	}

	fmt.Printf("Имя: %s, Возраст: %d\n", person.Name, person.Age)
}
