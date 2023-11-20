package main
import "fmt"
func main() {
	var first, second, third int
	fmt.Scanln(&first, &second, &third)
	if first < 0 || second < 0 || third < 0 {
		fmt.Println("Некорректный ввод")
	} else if first != second && second != third && third != first{
		fmt.Println("Все числа разные")
	} else if first == second && second == third {
		fmt.Println("Все числа равны")
	} else {
		fmt.Println("Два числа равны")
	}
}
