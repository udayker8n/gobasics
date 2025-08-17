package main

import (
	"errors"
	"time"
	"fmt"
)

type Account struct {
    accountNumber string
    balance       float64
    currency      string
	transcations []Transaction
}

type Transaction struct {
    ID          string
    Amount      float64
    Description string
    Type        string
    Timestamp   time.Time
}

// Method defined at package level, not inside main
func (ac Account) GetBalance() float64 {
    return ac.balance
}
func (ac Account) AccountNumber() string{
    return ac.accountNumber
}

func (ac Account) Currency() string{
    return ac.currency
}

func (ac *Account) Deposit(amount float64, description string) error {
    if amount <= 0 {
        return errors.New("deposit amount must be positive")
    }
    
    ac.balance += amount
    transaction := Transaction{
        ID:          generateTranscationID(),
        Amount:      amount,
        Type:        "DEPOSIT",
        Description: description,
        Timestamp:   time.Now(),
    }
    ac.transcations = append(ac.transcations, transaction)
    return nil
}
func (ba *Account) Withdraw(amount float64, description string) error {
    if amount <= 0 {
        return errors.New("withdrawal amount must be positive")
    }
    if ba.balance < amount {
        return errors.New("insufficient funds")
    }
    
    ba.balance -= amount
    transaction := Transaction{
        ID:          generateTranscationID(),
        Amount:      -amount,
        Type:        "WITHDRAWAL",
        Description: description,
        Timestamp:   time.Now(),
    }
    ba.transcations = append(ba.transcations, transaction)
    return nil
}
func generateTranscationID() string {
    return fmt.Sprintf("TXN-%d", time.Now().UnixNano())
}

func main() {
    // You can use Account and Transaction here
}
