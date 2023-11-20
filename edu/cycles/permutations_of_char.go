package main
import "fmt"
func main() {
	var num, dp, i int
	fmt.Scanln(&num)
	i = 1
	dp = 1
	for i <= num {
		dp *= i
		i++
	}
	fmt.Println(dp)
}
