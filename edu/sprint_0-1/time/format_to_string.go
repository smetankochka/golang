package main

import (
	"fmt"
	"time"
)

func FormatTimeToString(timestamp time.Time, format string) string {
	return timestamp.Format(format)
}

func main() {
	timestamp := time.Date(2023, 10, 23, 2, 41, 49, 0, time.UTC)
	format := "2006-01-02 15:04:05"
	result := FormatTimeToString(timestamp, format)
	fmt.Println(result) // Output: 2023-10-23 02:41:49
}
