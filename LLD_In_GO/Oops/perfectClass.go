package LLD

import (
	"fmt"
)

type BankAccount struct {
	accountNumber string
	ownerNumber   string
	balance       float64
}

// creating constructor in go :

/*
   Why NewBankAccount?
   By convention, Go uses:
   func NewUser(...) *BankAccount

   It can also return a value instead of a pointer:
   But pointers are commonly used when the object is intended to be modified or is relatively large.

*/

func NewBankAccount(accountNumber string, ownerName string) *BankAccount {
	// Initialize fields. Balance should start at 0.
	return &BankAccount{
		accountNumber: accountNumber,
		ownerNumber:   ownerName,
		balance:       0,
	}
}

// reciver :
func (b *BankAccount) Deposit(amount float64) {
	// Add amount to balance (only if amount is positive)

	if amount < 0 {
		fmt.Errorf("amount can't be negative")
	}

	b.balance = b.balance + amount
}

func (b *BankAccount) Withdraw(amount float64) bool {
	// Remove amount from balance if sufficient funds exist
	// Return true if successful, false otherwise

	if b.balance > amount {
		b.balance = b.balance - amount
		return true
	}

	return false
}

func (b *BankAccount) GetBalance() float64 {
	return b.balance
}

func main() {
	rudresh := NewBankAccount("59050100004223", "BARB0")

	fmt.Println(*rudresh) // well i am derefrencing as per cpp knowldge but go does it automatically

	rudresh.Deposit(100000) //pm

	fmt.Println(rudresh.GetBalance())

	rudresh.Withdraw(1000)

	fmt.Println(rudresh.GetBalance())

	fmt.Println(rudresh.balance)
}
