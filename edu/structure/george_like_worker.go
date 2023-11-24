package main

import (
	"fmt"
)

type Employee struct {
	name     string
	position string
	salary   int
	bonus    int
}

func (self Employee) CalculateTotalSalary() int {
	all := self.salary + self.bonus
	return all
}

func main() {
	gosha := Employee{"G", "adress", 1, 1}
	all := gosha.CalculateTotalSalary()
	fmt.Println(all)
}
