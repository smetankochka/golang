package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateTemperature(tempChan chan<- float64) {
	for {
		temp := rand.Float64()*80 - 30
		tempChan <- temp
		time.Sleep(2 * time.Second)
	}
}

func main() {
	tempChan := make(chan float64)
	go GenerateTemperature(tempChan)
	for {
		select {
		case temp := <-tempChan:
			fmt.Printf("Сгенерирована температура: %.2f °C\n", temp)
		}
	}
}
