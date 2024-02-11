package main

import (
	"fmt"
)

type Person struct {
	name    string
	age     int
	address string
}

func (self Person) Print() {
	fmt.Println("Name: " + self.name)
	fmt.Println("Age:", self.age)
	fmt.Println("Address: " + self.address)
}

func main() {
	gosha := Person{"G", 21, "adress"}
	gosha.Print()
}
