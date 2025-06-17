package main

import (
	"fmt"
)

func main() {
	fmt.Println("🏦 Welcome to Go Bank Management System")
	fmt.Println("--------------------------------------")
	newsvaccount := newSavingAccount("Nitin", "Delhi", "9876543210", "SAV123", 1000.0, 3.5)
	fmt.Println(newsvaccount)
	newcurrentaccount := newCurrentAccount("Nishant", "Patna", "7488204975", "CUR453", 2000.0, 4.5)
	fmt.Println(newcurrentaccount)
	//explicit declaration

	//var acc BankAccount // Declare interface variable
	//
	//acc = newSavingsAccount("Nitin", "Delhi", "9876543210", "SAV123", 1000.0, 3.5) // Assigning a struct instance that implements BankAccount
	//
	//acc.Deposit(200)
	//acc.Withdraw(100)
	//fmt.Println("Balance:", acc.GetBalance())
	//acc.PrintDetails()

}
