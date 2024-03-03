package main

import (
	"time"
)

func NextWorkday(start time.Time) time.Time {
	for {
		start = start.AddDate(0, 0, 1) // увеличиваем дату на 1 день

		// Проверяем, является ли текущая дата рабочим днем (Пн-Пт)
		if start.Weekday() != time.Saturday && start.Weekday() != time.Sunday {
			// Нашли следующий рабочий день, выходим из цикла и возвращаем его
			return start
		}
	}
}

func main() {
	start := time.Date(2024, 3, 3, 13, 0, 0, 0, time.UTC) // текущая дата
	nextWorkday := NextWorkday(start)
	println(nextWorkday) // выводим следующий рабочий день
}
