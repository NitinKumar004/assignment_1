package main

// Bankaccount interface declares operations all account must support
type BankAccount interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	GetBalance() float64
	PrintDetails()
}
type Customer struct {
	name    string
	Address string
	Phone   string
}
type SavingAccount struct {
	Customer
	AccountNumber string
	Balance       float64
	interestrate  float64
}
type CurrentAccount struct {
	Customer
	AccountNumber  string
	Balance        float64
	overdraftlimit float64
}

//constructor

func newSavingAccount(name, addr, phone, accno string, balance, interestrate float64) *SavingAccount {
	return &SavingAccount{Customer{name, addr, phone}, accno, balance, interestrate}
}
func newCurrentAccount(name, accno, phone, addr string, balance, ovd float64) *CurrentAccount {
	return &CurrentAccount{Customer{name, addr, phone}, accno, balance, ovd}
}
