package main

import "fmt"

func main() {
	var num, dp int
	dp = 0
	fmt.Scanln(&num)
	for i := 1; i <= num; i++ {
		dp *= 2
		dp += 1
	}
	fmt.Println(num * (num + 1) / 2)
}
