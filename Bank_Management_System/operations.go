package main

import (
	"errors"
	"fmt"
)

// saving account all functionalities
func (sa *SavingAccount) Deposit(amount float64) {
	sa.Balance += amount
}
func (sa *SavingAccount) Withdraw(amount float64) error {
	if sa.Balance <= amount {
		//fmt.Println("Insufficient funds")
		return errors.New("Insufficient funds")
	}
	sa.Balance -= amount
	fmt.Println("Balance withdrawn : ", amount)
	return nil
}
func (sa *SavingAccount) GetBalance() float64 {
	return sa.Balance

}
func (sa *SavingAccount) Printdetails() {
	fmt.Println("Saving Account details")

	fmt.Println("Balance : ", sa.Balance, "Name", sa.name, " Address : ", sa.Address, " Balance : ", sa.Balance, " Interest rate", sa.interestrate)
}

// current account all functionalities

func (ca *CurrentAccount) Deposit(amount float64) {
	ca.Balance += amount
}
func (ca *CurrentAccount) Withdraw(amount float64) error {
	if ca.Balance+ca.overdraftlimit <= amount {
		return errors.New("Insufficient funds")
	}
	ca.Balance -= amount
	return nil
}
func (ca *CurrentAccount) GetBalance() float64 {
	return ca.Balance
}
func (ca *CurrentAccount) Printdetails() {
	fmt.Println("Saving Account details")

	fmt.Println("Balance : ", ca.Balance, "Name", ca.name, " Address : ", ca.Address, " Balance : ", ca.Balance, " Interest rate", ca.overdraftlimit)
}
