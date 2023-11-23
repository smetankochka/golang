package main

import (
	"fmt"
	"lesson/students"
)

func main() {
	stud := students.Student{Name: "Kate", Year: 8, Gender: "M", Email: "dd.ru"}
	fmt.Println(stud)
}
