package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := Person{Name: "John", Age: 30}
	jsonBytes, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))
}
