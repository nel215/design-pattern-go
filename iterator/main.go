package main

import (
	"./book"
	"fmt"
)

func main() {
	bookShelf := book.NewBookShelf(4)
	bookShelf.AppendBook(book.NewBook("Around the World in 80 Days"))
	bookShelf.AppendBook(book.NewBook("Bible"))
	bookShelf.AppendBook(book.NewBook("Cinderella"))
	bookShelf.AppendBook(book.NewBook("Daddy-Long-Legs"))
	it := bookShelf.Iterator()
	for it.HasNext() {
		book, _ := it.Next().(*book.Book)
		fmt.Println(book.GetName())
	}
}
