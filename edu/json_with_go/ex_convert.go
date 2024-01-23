package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name         string `json:"name"` // `json:"name"` — тег поля. Без него в JSON будет ключ "Name" вместо "name".
	Age          int    `json:"age"`
	Gender       string `json:"-"` // поле с тегом `json:"-"` при кодировании json игнорируется
	privateNotes string // неэкспортируемые поля так же игнорируются
}

func main() {
	jsonStr := `{"name": "John", "age": 30, "Gender": "male"}`
	var person Person
	err := json.Unmarshal([]byte(jsonStr), &person)
	if err != nil {
		panic(err)
	}
	fmt.Println(person)
}
