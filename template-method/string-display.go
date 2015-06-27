package main

import "fmt"

type StringDisplay struct {
	string string
}

func (self *StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < len(self.string); i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func (self *StringDisplay) Open() {
	self.printLine()
}

func (self *StringDisplay) Print() {
	fmt.Println("|" + self.string + "|")
}

func (self *StringDisplay) Close() {
	self.printLine()
}
