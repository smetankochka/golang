package main

import (
	"fmt"
	"time"
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Report struct {
	User
	ReportID int
	Date     string
}

func (u User) PrintInfo() {
	fmt.Printf("User ID: %d\nName: %s\nEmail: %s\nAge: %d\n", u.ID, u.Name, u.Email, u.Age)
	fmt.Println("----------------------------")
}

func (r Report) PrintInfo() {
	fmt.Println("Report Information:")
	r.User.PrintInfo()
	fmt.Printf("Report ID: %d\nDate: %s\n", r.ReportID, r.Date)
	fmt.Println("----------------------------")
}

func CreateReport(user User, reportDate string) Report {
	reportID := int(time.Now().Unix()) // Генерация уникального ID на основе времени
	return Report{User: user, ReportID: reportID, Date: reportDate}
}

func GenerateUserReports(users []User, reportDate string) []Report {
	var reports []Report

	for _, user := range users {
		report := CreateReport(user, reportDate)
		reports = append(reports, report)
	}

	return reports
}

func main() {
	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 30},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Age: 25},
	}

	reportDate := "2024-02-07"
	userReports := GenerateUserReports(users, reportDate)

	for _, report := range userReports {
		report.PrintInfo()
	}
}
