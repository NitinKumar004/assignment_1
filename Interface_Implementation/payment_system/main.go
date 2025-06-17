package main

import (
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float64) string
}

//extra interface to otp sending
type Generateotp interface {
	Generateotp()
}
type CreditCard struct {
	CardNumber string
}
type Paypal struct {
	Email string
}
type UPI struct {
	upiid string
}

//only add otp system to pay by credit card
func (c CreditCard) Generateotp() {
	fmt.Println("[CreditCard] OTP sent to registered number")
}
func (c CreditCard) Pay(amount float64) string {
	last4 := "XXXX"
	if len(c.CardNumber) >= 4 {
		last4 = c.CardNumber[len(c.CardNumber)-4:]
	}
	return fmt.Sprintf("[CreditCard] Paid ₹%.2f using card ending with %s", amount, last4)
}
func (p Paypal) Pay(amount float64) string {
	return fmt.Sprintf("[PayPal] Paid ₹%.2f using PayPal account: %s", amount, p.Email)
}
func (u UPI) Pay(amount float64) string {
	return fmt.Sprintf("[UPI] Paid ₹%.2f using UPI: %s", amount, u.upiid)
}

func main() {
	//wrapp all instance in single to loop through
	payments := []PaymentMethod{
		CreditCard{CardNumber: "1234567812341234"},
		Paypal{Email: "nitin@GMAIL.COM"},
		UPI{upiid: "nitin@ybl"},
	}
	amount := 500.0
	//We Use Type Assertion
	for _, payment := range payments {
		if otpmethod, ok := payment.(Generateotp); ok {
			otpmethod.Generateotp()

		}
		fmt.Println(payment.Pay(amount))
		fmt.Println()
	}

}
