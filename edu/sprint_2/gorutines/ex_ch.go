package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 123 // отправляем значение в канал
	}()

	val := <-ch      // получаем значение из канала
	fmt.Println(val) // выводит "123"
}
