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
// Base account information
type BaseAccount struct {
    ID       uint64
    Owner    string
    Currency string
    Created  time.Time
}

func (ba BaseAccount) GetAccountInfo() string {
    return fmt.Sprintf("Account %d owned by %s in %s", 
                      ba.ID, ba.Owner, ba.Currency)
}

// Savings account embeds BaseAccount
type SavingsAccount struct {
    BaseAccount              // Embedding - fields and methods are promoted
    Balance        float64
    InterestRate   float64
    MinimumBalance float64
}

func (sa *SavingsAccount) CalculateInterest() float64 {
    return sa.Balance * sa.InterestRate / 100
}

func (sa *SavingsAccount) ApplyMonthlyInterest() {
    interest := sa.CalculateInterest()
    sa.Balance += interest
    fmt.Printf("Applied interest: $%.2f, New balance: $%.2f\n", 
               interest, sa.Balance)
}

// Investment account with different behavior
type InvestmentAccount struct {
    BaseAccount                // Same embedding
    Holdings    []StockHolding
    TotalValue  float64
    RiskLevel   string
}

type StockHolding struct {
    Symbol       string
    Shares       int
    Price        float64
    PurchaseDate time.Time
}

func (ia *InvestmentAccount) AddHolding(symbol string, shares int, price float64) {
    holding := StockHolding{
        Symbol:       symbol,
        Shares:       shares,
        Price:        price,
        PurchaseDate: time.Now(),
    }
    ia.Holdings = append(ia.Holdings, holding)
    ia.updateTotalValue()
}

func (ia *InvestmentAccount) updateTotalValue() {
    total := 0.0
    for _, holding := range ia.Holdings {
        total += float64(holding.Shares) * holding.Price
    }
    ia.TotalValue = total
}


func main() {
    // You can use Account and Transaction here
}
