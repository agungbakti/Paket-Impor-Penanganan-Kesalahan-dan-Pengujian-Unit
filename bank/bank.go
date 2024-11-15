package bank

import (
    "errors"
    "fmt"
)

type Account struct {
    balance float64
}

func (a *Account) Deposit(amount float64) error {
	//jika amount kurang dari atau sama dengan 0, return error
    if amount <= 0 {
        return errors.New("amount must be positive")
    }
	//menambahkan amount ke balance
    a.balance += amount
    return nil
}

func (a *Account) Withdraw(amount float64) error {
	//jika amount kurang dari atau sama dengan 0, return error
    if amount <= 0 {
        return fmt.Errorf("amount must be positive %.1f", amount)
    }
    
	//jika amount lebih besar dari balance
    if amount > a.balance {
        return fmt.Errorf("cannot withdraw %.1f from balance of %.1f", amount, a.balance)
    }
    
	//kurangi amount dari balance
    a.balance -= amount
    return nil
}

func (a *Account) GetBalance() float64 {
	//mengembalikan balance
    return a.balance
}