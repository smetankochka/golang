package main
import "fmt"
func main() {
	var num int
	fmt.Scanln(&num)
	if num > 0 {
		fmt.Println("Число положительное")
	} else {
		fmt.Println("Число отрицательное")
	}
}
