package main

import (
	"fmt"
	"sort"
)

func SortNames(names []string) {
	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})
}

func main() {
	names := []string{"Арина", "Варвара", "Есения", "Аксинья"}

	fmt.Println("Before sorting:")
	for _, name := range names {
		fmt.Println(name)
	}

	SortNames(names)

	fmt.Println("\nAfter sorting:")
	for _, name := range names {
		fmt.Println(name)
	}
}
