package main

import "fmt"

type CharDispaly struct {
	ch string
}

func (self *CharDispaly) Open() {
	fmt.Print("<<")
}

func (self *CharDispaly) Print() {
	fmt.Print(self.ch)
}

func (self *CharDispaly) Close() {
	fmt.Println(">>")
}
