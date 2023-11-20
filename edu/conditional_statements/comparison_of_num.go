package main
import "fmt"
func main() {
	var first, second int
	fmt.Scanln(&first, &second)
	if first > second {
		fmt.Println("Первое число больше второго")
	} else if second > first{
		fmt.Println("Второе число больше первого")
	} else {
		fmt.Println("Числа равны")
	}
}
