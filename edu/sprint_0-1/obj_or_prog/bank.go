package main

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance float64
}

func NewAccount(owner string) *Account {
	return &Account{owner: owner, balance: 0}
}

func (acc *Account) SetBalance(amount float64) error {
	if amount < 0 {
		return errors.New("Баланс не может быть отрицательным")
	}
	acc.balance = amount
	return nil
}

func (acc *Account) GetBalance() float64 {
	return acc.balance
}

func (acc *Account) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("Нельзя внести отрицательную сумму")
	}
	acc.balance += amount
	return nil
}

func (acc *Account) Withdraw(amount float64) error {
	if amount < 0 {
		return errors.New("Нельзя снять отрицательную сумму")
	}
	if acc.balance < amount {
		return errors.New("Недостаточно средств на счете")
	}
	acc.balance -= amount
	return nil
}

func main() {
	// Создаем новый аккаунт для пользователя Alice
	aliceAcc := NewAccount("Alice")

	// Устанавливаем начальный баланс
	err := aliceAcc.SetBalance(1000)
	if err != nil {
		fmt.Println(err)
	}

	// Информация о балансе аккаунта Алисы
	fmt.Printf("Текущий баланс пользователя %s: %.2f\n", aliceAcc.owner, aliceAcc.GetBalance())

	// Депозит денег на счет
	err = aliceAcc.Deposit(500)
	if err != nil {
		fmt.Println(err)
	}

	// Информация о балансе после депозита
	fmt.Printf("После внесения депозита. Текущий баланс пользователя %s: %.2f\n", aliceAcc.owner, aliceAcc.GetBalance())

	// Снятие денег со счета
	err = aliceAcc.Withdraw(200)
	if err != nil {
		fmt.Println(err)
	}

	// Информация о балансе после снятия средств
	fmt.Printf("После снятия денег. Текущий баланс пользователя %s: %.2f\n", aliceAcc.owner, aliceAcc.GetBalance())
}
