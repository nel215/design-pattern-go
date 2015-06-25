package book

// Book
type Book struct {
	name string
}

func (self *Book) Book(name string) {
	self.name = name
}

func (self *Book) GetName() string {
	return self.name
}

func NewBook(name string) *Book {
	return &Book{name}
}
