package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b BankAccount) DisplayBalance() {
	fmt.Println("Balance:", b.Balance)
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
	fmt.Println("you have successfully deposited", amount)
	fmt.Println("New balance:", b.Balance)

}

func (b *BankAccount) Withdraw(amount float64) {
	if amount > b.Balance {
		fmt.Println("You amount is greater than you balance")
		b.DisplayBalance()
		return
	}
	b.Balance -= amount
	fmt.Println("You have successfully withdrawn", amount)
}
func main() {
	b := BankAccount{"nitin", 4000.0}
	b.DisplayBalance()
	b.Deposit(400)
	b.Withdraw(400)
	b.DisplayBalance()

}
