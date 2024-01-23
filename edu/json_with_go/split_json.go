package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Student struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}

func splitJSONByClass(jsonData []byte) (map[string][]byte, error) {
	jsonstr := string(jsonData)
	anstmp := map[string][]Student{}
	students := []Student{}
	reader := strings.NewReader(jsonstr)
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&students)
	if err != nil {
		return nil, err
	}
	for _, st := range students {
		anstmp[st.Class] = append(anstmp[st.Class], st)
	}
	ans := map[string][]byte{}
	for key, value := range anstmp {
		jsonbytes, err2 := json.Marshal(value)
		if err2 != nil {
			return nil, err2
		}
		ans[key] = jsonbytes
	}
	return ans, nil
}

func main() {
	jsonstr := `[
		{
		"name": "Oleg",
		"class": "9B"
		},
		{
		"name": "Ivan",
		"class": "9A"
		},
		{
		"name": "Maria",
		"class": "9B"
		},
		{
		"name": "John",
		"class": "9A"
		}
	]`
	ans, err := splitJSONByClass([]byte(jsonstr))
	if err != nil {
		panic(err)
	}
	for k, v := range ans {
		fmt.Println(k, string(v))
	}
}
