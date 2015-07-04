package main

import (
	"./singleton"
	"fmt"
)

func main() {
	obj1 := singleton.GetInstance()
	obj2 := singleton.GetInstance()
	if obj1 == obj2 {
		fmt.Println("ok")
	} else {
		panic("ng")
	}
	obj1.Print()
}
