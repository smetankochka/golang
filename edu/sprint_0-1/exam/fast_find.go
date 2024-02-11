package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Сначала читаем количество людей в списке
	scanner.Scan()
	var n int
	fmt.Sscan(scanner.Text(), &n)

	// Затем читаем и составляем список людей
	names := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		names[i] = scanner.Text()
	}

	// Сортируем список по алфавиту
	sort.Strings(names)

	// Читаем префиксы для поиска
	var prefixes []string
	for scanner.Scan() {
		prefix := scanner.Text()
		if prefix == "" {
			break
		}
		prefixes = append(prefixes, prefix)
	}

	// Производим поиск и выводим результат
	for _, prefix := range prefixes {
		found := false
		for _, name := range names {
			if len(name) >= len(prefix) && name[:len(prefix)] == prefix {
				fmt.Println(name)
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Не найдено")
		}
	}
}
