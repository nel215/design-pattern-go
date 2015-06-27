package main

import "fmt"

type Banner struct {
	string string
}

func NewBanner(string string) *Banner {
	return &Banner{string}
}

func (self *Banner) ShowWithParen() {
	fmt.Println("(" + self.string + ")")
}

func (self *Banner) ShowWithAster() {
	fmt.Println("*" + self.string + "*")
}
