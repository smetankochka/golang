package main

import (
	"log"
)

// Order представляет информацию о заказе.
type Order struct {
	OrderNumber  int
	CustomerName string
	OrderAmount  float64
}

// OrderLogger представляет журнал заказов и хранит записи о заказах.
type OrderLogger struct{}

// NewOrderLogger создает новый экземпляр OrderLogger.
func NewOrderLogger() *OrderLogger {
	return &OrderLogger{}
}

// AddOrder записывает информацию о добавленном заказе в терминал.
func (logger *OrderLogger) AddOrder(order Order) {
	log.Printf("Добавлен заказ #%d, Имя клиента: %s, Сумма заказа: $%.2f\n", order.OrderNumber, order.CustomerName, order.OrderAmount)
}

func main() {
	orderLogger := NewOrderLogger()

	// Пример добавления заказа
	order := Order{OrderNumber: 1234, CustomerName: "Иванов", OrderAmount: 100.5}
	orderLogger.AddOrder(order)
}
