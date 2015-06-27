package main

func main() {
	var p Print = NewPrintBanner("hello")
	p.PrintWeak()
	p.PrintStrong()
}
