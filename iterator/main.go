package main

import (
	"fmt"
)

func main() {
	bookShelf := NewBookShelf(4)
	bookShelf.AppendBook(NewBook("Around the World in 80 Days"))
	bookShelf.AppendBook(NewBook("Bible"))
	bookShelf.AppendBook(NewBook("Cinderella"))
	bookShelf.AppendBook(NewBook("Daddy-Long-Legs"))
	it := bookShelf.Iterator()
	for it.HasNext() {
		book, _ := it.Next().(*Book)
		fmt.Println(book.GetName())
	}
}
