package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
	"time"
)

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	file, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var logsInRange []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line) // Разбиваем строку на подстроки

		dateString := parts[0]
		logDate, err := time.Parse("02.01.2006", dateString) // Парсим дату строки
		if err != nil {
			return nil, err
		}

		if (logDate.After(start) && logDate.Before(end)) || logDate == start || logDate == end {
			logsInRange = append(logsInRange, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(logsInRange) == 0 {
		return nil, errors.New("No log entries found in the specified range")
	}

	return logsInRange, nil
}

func main() {
	startDate, _ := time.Parse("02.01.2006", "13.12.2022")
	endDate, _ := time.Parse("02.01.2006", "15.12.2022")

	logs, err := ExtractLog("yourLogFile.txt", startDate, endDate)
	if err != nil {
		panic(err)
	}

	for _, log := range logs {
		println(log)
	}
}
