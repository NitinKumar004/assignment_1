package main

import "fmt"

type Logger interface {
	Log()
	Save()
}
type File struct {
}

func (f File) Log() {
	fmt.Println("Logging from value receiver")
}
func (f File) Save() {
	fmt.Println("Saving from pointer receiver")
}

//	func (f File) Save() {
//		fmt.Println("Saving from pointer receiver")
//	}
func main() {
	var l Logger
	f1 := File{}
	f2 := &File{}
	l = f1 // Try this
	l.Log()
	l.Save()
	//l.Save()
	l = f2 // Try this too
	l.Log()
	l.Save()
}
