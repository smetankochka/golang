package main

import (
	"fmt"
	"sort"
)

type Company struct {
	Workers []Worker
}

type Worker struct {
	Name       string
	Position   string
	Salary     uint
	Experience uint
}

func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) {
	worker := Worker{
		Name:       name,
		Position:   position,
		Salary:     salary,
		Experience: experience,
	}
	c.Workers = append(c.Workers, worker)
}

func (c Company) SortWorkers() ([]string, error) {
	sort.Slice(c.Workers, func(i, j int) bool {
		if c.Workers[i].Salary*12*c.Workers[i].Experience != c.Workers[j].Salary*12*c.Workers[j].Experience {
			return c.Workers[i].Salary*12*c.Workers[i].Experience > c.Workers[j].Salary*12*c.Workers[j].Experience
		}
		return getPositionOrder(c.Workers[i].Position) > getPositionOrder(c.Workers[j].Position)
	})
	sortedData := make([]string, len(c.Workers))
	for i, worker := range c.Workers {
		sortedData[i] = fmt.Sprintf("%s - %d - %s", worker.Name, worker.Salary*12*worker.Experience, worker.Position)
	}

	return sortedData, nil
}

func getPositionOrder(position string) int {
	switch position {
	case "директор":
		return 5
	case "зам. директора":
		return 4
	case "начальник цеха":
		return 3
	case "мастер":
		return 2
	case "рабочий":
		return 1
	default:
		return 0
	}
}

func main() {
	company := Company{}

	// Пример использования
	company.AddWorkerInfo("Михаил", "директор", 12000, 24)
	company.AddWorkerInfo("Андрей", "мастер", 11800, 12)
	company.AddWorkerInfo("Игорь", "зам. директора", 11000, 36)

	sortedData, err := company.SortWorkers()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	for _, data := range sortedData {
		fmt.Println(data)
	}
}
