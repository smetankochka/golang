package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Student struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func modifyJSON(jsonData []byte) ([]byte, error) {
	jsonstr := string(jsonData)
	students := []Student{}
	decoder := json.NewDecoder(strings.NewReader(jsonstr))
	err := decoder.Decode(&students)
	if err != nil {
		return nil, err
	}
	length := len(students)
	for i := 0; i < length; i++ {
		students[i].Grade++
	}
	jsonBytes, err2 := json.Marshal(students)
	if err2 != nil {
		return nil, err2
	}
	return jsonBytes, nil
}

func main() {
	jsonstr := `[
		{
			"name": "Oleg",
			"grade": 9
		},
		{
			"name": "Katya",
			"grade": 10
		}
	]`
	ans, err := modifyJSON([]byte(jsonstr))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(string(ans))
}
