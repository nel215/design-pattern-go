package main

type PrintBanner struct {
	banner *Banner
}

func NewPrintBanner(string string) *PrintBanner {
	return &PrintBanner{NewBanner(string)}
}

func (self *PrintBanner) PrintWeak() {
	self.banner.ShowWithParen()
}

func (self *PrintBanner) PrintStrong() {
	self.banner.ShowWithAster()
}
