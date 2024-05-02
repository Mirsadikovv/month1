package main

import "fmt"

type Book struct {
	Title     string
	Author    string
	ISBN      int
	Available bool
}

type Library struct {
	Books []Book
}

func (b *Book) BorrowBook() string {
	if b.Available {
		b.Available = false
		return "successefull borrowed"
	}
	return "this book is already taken"

}

func (b *Book) ReturnBook() string {
	b.Available = true
	return "successefull returned"
}

func (b Book) PrintDetails() {
	fmt.Println(b)
}

func (lib *Library) AddBook(book Book) {
	lib.Books = append(lib.Books, book)
	fmt.Println(book, " successefully added to library")
}

func (lib *Library) FindBookByISBN(isbn int) *Book {
	for i, book := range lib.Books {
		if book.ISBN == isbn {
			return &lib.Books[i]
		}
	}
	return nil
}

func (lib *Library) ListAvailableBooks() (availableBooks []Book) {
	for _, book := range lib.Books {
		if book.Available {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func main() {
	book1 := Book{Title: "Onegin", Author: "pushkin", ISBN: 1233, Available: true}
	book2 := Book{Title: "mertvie dushi", Author: "gogol", ISBN: 1313, Available: true}

	library := Library{}
	library.AddBook(book1)
	library.AddBook(book2)
	book1.BorrowBook()
	book1.PrintDetails()

	book1.ReturnBook()
	book1.PrintDetails()

	book1.BorrowBook()
	book1.PrintDetails()
	fmt.Println(*library.FindBookByISBN(1233))
	// library.ListAvailableBooks()
}
