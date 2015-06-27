package main

func main() {
	charDispaly := &Display{&CharDispaly{"a"}}
	charDispaly.dispaly()
	stringDispaly := &Display{&StringDisplay{"Hello, world!"}}
	stringDispaly.dispaly()
}
