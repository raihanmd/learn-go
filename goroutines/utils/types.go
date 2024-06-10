package utils

import (
	"fmt"
	"sync"
)

type DbResult struct {
	Id, Name string
}

type BankAccount struct {
	sync.RWMutex
	Name    string
	Balance int
}

func (account *BankAccount) Lock() {
	account.RWMutex.Lock()
}

func (account *BankAccount) Unlock() {
	account.RWMutex.Unlock()
}

func (account *BankAccount) Deposit(amount int) {
	account.Balance += amount
}

func Transfer(from, to *BankAccount, amount int) {
	from.Lock()
	fmt.Println("Lock user1", from.Name)
	defer from.Unlock()
	to.Lock()
	fmt.Println("Lock user2", to.Name)
	defer to.Unlock()
	from.Deposit(-amount)
	to.Deposit(amount)
}

type DbService interface {
	GetDbResult() []DbResult
}
