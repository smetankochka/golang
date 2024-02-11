package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Student struct {
	Name          string    `json:"name"`
	AdmissionDate time.Time `json:"admission_date"`
	DaysOnCourse  int       `json:"-"`
}

func processJSON(jsonData []byte) error {
	var students []Student
	if err := json.Unmarshal(jsonData, &students); err != nil {
		return err
	}
	currentDate, err := time.Parse("2006-01-02", "2023-10-01")
	if err != nil {
		return err
	}
	for _, student := range students {
		daysOnCourse := int(currentDate.Sub(student.AdmissionDate).Hours() / 24)
		fmt.Printf("%s: %d\n", student.Name, daysOnCourse)
	}
	return nil
}

func main() {
	jsonstr := `[
		{
		"name": "Анна",
		"admission_date": "2023-09-29T00:00:00Z"
		},
		{
		"name": "Иван",
		"admission_date": "2023-09-28T00:00:00Z"
		}
	]`
	err := processJSON([]byte(jsonstr))
	if err != nil {
		panic(err)
	}
}
