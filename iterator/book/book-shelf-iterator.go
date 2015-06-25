package book

// BookShelfIterator
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func (self *BookShelfIterator) HasNext() bool {
	if self.index < self.bookShelf.getLength() {
		return true
	} else {
		return false
	}
}

func (self *BookShelfIterator) Next() interface{} {
	book := self.bookShelf.getBookAt(self.index)
	self.index++
	return book
}
