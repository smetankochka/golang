package main
import "fmt"
func main() {
	var num int
	fmt.Scanln(&num)
	if num < 0 {
		fmt.Println("Некорректный ввод")
	} else if num < 10 {
		fmt.Println("Число меньше 10")
	} else if num < 100 {
		fmt.Println("Число меньше 100")
	} else if num < 1000 {
		fmt.Println("Число меньше 1000")
	} else{
		fmt.Println("Число больше или равно 1000")
	}
}
