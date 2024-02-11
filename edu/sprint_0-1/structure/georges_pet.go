package main

import (
	"fmt"
)

type Animal interface {
	MakeSound()
}

type Dog struct {
}

type Cat struct {
}

func (self Dog) MakeSound() {
	fmt.Println("Гав!")
}

func (self Cat) MakeSound() {
	fmt.Println("Мяу!")
}

func main() {
	var c Animal = Cat{}
	var d Animal = Dog{}
	c.MakeSound()
	d.MakeSound()
}
