package students

import "fmt"

type Student struct {
	Name     string
	Year     int
	Gender   string
	Email    string
	password string
	snils    string
}

func (self Student) Printdata(s string) {
	if s == "Name" {
		fmt.Println(self.Name)
	} else if s == "Year" {
		fmt.Println(self.Year)
	} else if s == "Gender" {
		fmt.Println(self.Gender)
	} else if s == "Email" {
		fmt.Println(self.Email)
	} else {
		fmt.Println("nothing about this")
	}
}
