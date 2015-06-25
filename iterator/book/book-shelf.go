package book

// BookShelf
type BookShelf struct {
	books []*Book
	last  int
}

func (self *BookShelf) getBookAt(index int) *Book {
	return self.books[index]
}

func (self *BookShelf) AppendBook(book *Book) {
	self.books[self.last] = book
	self.last++
}

func (self *BookShelf) getLength() int {
	return self.last
}

func NewBookShelf(maxsize int) *BookShelf {
	return &BookShelf{make([]*Book, maxsize), 0}
}

func (self *BookShelf) Iterator() Iterator {
	return &BookShelfIterator{self, 0}
}
