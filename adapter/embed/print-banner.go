package main

type PrintBanner struct {
	*Banner
}

func NewPrintBanner(string string) *PrintBanner {
	return &PrintBanner{NewBanner(string)}
}

func (self *PrintBanner) PrintWeak() {
	self.ShowWithParen()
}

func (self *PrintBanner) PrintStrong() {
	self.ShowWithAster()
}
