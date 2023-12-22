package main

import "fmt"

func countWaysToReachFloor(n int) int {
	dp := make([]int, n+3)
	dp[n] = 1

	for i := n - 1; i >= 0; i-- {
		dp[i] = dp[i+1] + dp[i+2] + dp[i+3] // Суммируем способы для i+1, i+2 и i+3
	}

	return dp[0]
}

func main() {
	ways := countWaysToReachFloor(5)
	fmt.Println(ways)
}
