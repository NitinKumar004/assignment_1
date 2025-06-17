📘 Go Assignment: Pointers and Methods
🧩 Part 1: Understanding Pointers and Methods
Objective:
Understand how pointers and methods work in Go, including how method receivers (value vs pointer) affect behavior.
Task:
Implement a simple Bank Account system.
Instructions:
Define a BankAccount struct with the following fields:
Owner string
Balance float64
Implement the following methods:

func (b BankAccount) DisplayBalance()
func (b *BankAccount) Deposit(amount float64)
func (b *BankAccount) Withdraw(amount float64)
💡 DisplayBalance should print the owner and current balance.
💡 Deposit should increase the balance.
💡 Withdraw should decrease the balance, but only if sufficient funds exist.



In the main() function, do the following:


Create a BankAccount value.


Call each method and observe how balance changes or remains unchanged based on pointer/value receivers.


Expected Output:
A log of deposits and withdrawals, showcasing that value receiver doesn't change the struct unless a pointer is used.


🧪 Part 2: Try It Yourself – Method Sets in Go
Objective:
Understand how method sets work in Go by implementing and observing what compiles and what doesn’t.
Task:
Copy and paste the following code:

package main
import "fmt"
type Logger interface {
Log()
}
type File struct{}
func (f File) Log() {
fmt.Println("Logging from value receiver")
}
func (f *File) Save() {
fmt.Println("Saving from pointer receiver")
}
func main() {
var l Logger
f1 := File{}
f2 := &File{}
l = f1 // Try this
l.Log()
l = f2 // Try this too
l.Log()
}
Play around:


Comment/uncomment l = f1 and l = f2 one at a time.


Observe which assignments compile and which don't.


Try removing the Log() method from File and adding it only to *File. What changes?


🔍 What to Observe:
Which types satisfy the Logger interface?


How does defining a method on a pointer receiver affect interface satisfaction?
