package main

import (
	"fmt"
)

type UnderlinePen struct {
	decostr string
}

func NewUnderlinePen(decostr string) *UnderlinePen {
	return &UnderlinePen{decostr}
}

func (self *UnderlinePen) use(s string) {
	length := len(s)
	fmt.Println("\"" + s + "\"")
	fmt.Print(" ")
	for i := 0; i < length; i++ {
		fmt.Print(self.decostr)
	}
	fmt.Println()
}

func (self *UnderlinePen) createClone() Product {
	clone := *self
	return &clone
}
