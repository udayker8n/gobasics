// here you'll see how i learn interfaces
package main
 import (
	"fmt"
	"errors"
 )
// PaymentProcessor defines the contract for processing payments
type PaymentProcessor interface {
    ProcessPayment(amount float64) error
    GetName() string
}

// Credit card implementation
type CreditCardProcessor struct {
    CardNumber string
    CVV        string
}

func (c CreditCardProcessor) ProcessPayment(amount float64) error {
    fmt.Printf("Processing $%.2f via Credit Card ending in %s\n", 
               amount, c.CardNumber[len(c.CardNumber)-4:])
    
    // Simulate API call to payment gateway
    if amount > 10000 {
        return errors.New("amount exceeds credit card limit")
    }
    
    return nil
}

func (c CreditCardProcessor) GetName() string {
    return "Credit Card"
}

// PayPal implementation
type PayPalProcessor struct {
    Email string
}

func (p PayPalProcessor) ProcessPayment(amount float64) error {
    fmt.Printf("Processing $%.2f via PayPal account %s\n", amount, p.Email)
    
    // Simulate PayPal API integration
    if amount > 5000 {
        return errors.New("PayPal requires additional verification for large amounts")
    }
    
    return nil
}

func (p PayPalProcessor) GetName() string {
    return "PayPal"
}

// Using polymorphism with interfaces
func processPayment(processor PaymentProcessor, amount float64) {
    fmt.Printf("Using %s processor:\n", processor.GetName())
    err := processor.ProcessPayment(amount)
    if err != nil {
        fmt.Printf("Payment failed: %v\n", err)
    } else {
        fmt.Printf("Payment successful!\n")
    }
}



func main(){

}