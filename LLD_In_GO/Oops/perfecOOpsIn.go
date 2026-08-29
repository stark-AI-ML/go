package main

import "fmt"

type BankAccount struct {
	accountNumber string
	ownerName     string
	balance       float64
}

func NewBankAccount(accountNumber, ownerName string) *BankAccount {
	return &BankAccount{
		accountNumber: accountNumber,
		ownerName:     ownerName,
		balance:       0,
	}
}

func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.balance += amount
	}
}

func (b *BankAccount) Withdraw(amount float64) bool {
	if amount > 0 && b.balance >= amount {
		b.balance -= amount
		return true
	}
	return false
}

func (b *BankAccount) GetBalance() float64 {
	return b.balance
}

func main() {
	account := NewBankAccount("123456", "John Doe")
	account.Deposit(1000)
	fmt.Printf("%.1f\n", account.GetBalance())

	success := account.Withdraw(500)
	fmt.Println(success)
	fmt.Printf("%.1f\n", account.GetBalance())

	success = account.Withdraw(1000)
	fmt.Println(success)
}
