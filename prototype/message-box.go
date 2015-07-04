package main

import (
	"fmt"
)

type MessageBox struct {
	decostr string
}

func NewMessageBox(decostr string) *MessageBox {
	return &MessageBox{decostr}
}

func (self *MessageBox) use(s string) {
	length := len(s)
	for i := 0; i < length+4; i++ {
		fmt.Print(self.decostr)
	}
	fmt.Println()
	fmt.Println(self.decostr + " " + s + " " + self.decostr)
	for i := 0; i < length+4; i++ {
		fmt.Print(self.decostr)
	}
	fmt.Println()
}

func (self *MessageBox) createClone() Product {
	clone := *self
	return &clone
}
