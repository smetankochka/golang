package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n, m, k, p int
		fmt.Scan(&n, &m, &k, &p)

		rowLimits := make([][3]int, p) // Двумерный массив для ограничений на ряды
		colLimits := make([][3]int, p) // Двумерный массив для ограничений на места

		for j := 0; j < p; j++ { // Считываем ограничения
			var limitType string
			fmt.Scan(&limitType)

			if limitType == "R" {
				fmt.Scan(&rowLimits[j][0], &rowLimits[j][1], &rowLimits[j][2])
			} else if limitType == "C" {
				fmt.Scan(&colLimits[j][0], &colLimits[j][1], &colLimits[j][2])
			}
		}

		var q, disatisfaction int
		disatisfaction = 0
		fmt.Scan(&q)

		occupied := make([][]bool, n+1)
		for i := 1; i <= n; i++ {
			occupied[i] = make([]bool, m+1)
		}

		for j := 0; j < q; j++ { // Обрабатываем поручения
			var s, r, c int
			fmt.Scan(&s, &r, &c)

			if !occupied[r][c] {
				occupied[r][c] = true
			} else {
				disatisfaction++
			}

			for l := 0; l < p; l++ { // Проверяем ограничения
				if rowLimits[l][0] == s && (r < rowLimits[l][1] || r > rowLimits[l][2]) {
					disatisfaction++
					break
				}
				if colLimits[l][0] == s && (c < colLimits[l][1] || c > colLimits[l][2]) {
					disatisfaction++
					break
				}
			}
		}

		fmt.Println(k - q - disatisfaction)
	}
}
