package main

type Product interface {
	use(s string)
	createClone() Product
}
