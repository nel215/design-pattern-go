package main

type Display struct {
	strategy AbstractDisplay
}

func (self *Display) dispaly() {
	self.strategy.Open()
	for i := 0; i < 5; i++ {
		self.strategy.Print()
	}
	self.strategy.Close()
}
