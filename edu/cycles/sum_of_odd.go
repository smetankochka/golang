package main
import "fmt"
func main() {
	var num int
	fmt.Scanln(&num)
	if num < 0 {
		fmt.Println("Некорректный ввод")
		return
	}
	var sum int
	sum = 0
	for i := 1; i <= num; i += 2 {
		sum += i
	}
	fmt.Println(sum)
}
