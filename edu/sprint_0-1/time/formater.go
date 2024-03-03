package main

import (
	"fmt"
	"time"
)

func TimeAgo(pastTime time.Time) string {
	now := time.Now()
	diff := now.Sub(pastTime)

	years := diff.Hours() / (24 * 365)
	months := diff.Hours() / (24 * 30)
	days := diff.Hours() / 24
	hours := diff.Hours()
	minutes := diff.Minutes()

	if years >= 1 {
		return fmt.Sprintf("%.0f years ago", years)
	} else if months >= 1 {
		return fmt.Sprintf("%.0f months ago", months)
	} else if days >= 1 {
		return fmt.Sprintf("%.0f days ago", days)
	} else if hours >= 1 {
		return fmt.Sprintf("%.0f hours ago", hours)
	} else {
		return fmt.Sprintf("%.0f minutes ago", minutes)
	}
}

func main() {
	pastTime := time.Date(2023, 10, 23, 2, 41, 49, 0, time.UTC)
	result := TimeAgo(pastTime)
	fmt.Println(result) // Output: 1 month ago

	// Пример работы с часами
	pastTime2 := time.Now().Add(-2 * time.Hour)
	timeAgo := TimeAgo(pastTime2)
	fmt.Println(timeAgo) // Output: 2 hours ago
}
