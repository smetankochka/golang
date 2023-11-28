package main

import "fmt"

func main() {
	fmt.Println("Start")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	panic("Something went wrong")
	fmt.Println("End")
}
