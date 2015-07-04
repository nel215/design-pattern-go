package singleton

import "fmt"

type singleton struct {
}

var instance *singleton = &singleton{}

func GetInstance() *singleton {
	return instance
}

func (self *singleton) Print() {
	fmt.Println("hello")
}
