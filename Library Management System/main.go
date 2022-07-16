package main

import (
	"bufio"
	"fmt"
	"os"
)

type BookType int64

const (
	eBook BookType = iota
	Audiobook
	Hardback
	Paperback
	Encyclopedia
	Magazine
	Comic
)

type Book interface {
	Booktype() BookType
	Name() string
	Author() string
	// Borrow accepts a username and attempts to borrow the book in that user's name.
	// Returns a boolean indicating the success of the borrow
	Borrow(string) bool
	Return(string)
}

type PhysicalBook struct {
	bookType BookType
	name     string
	author   string
	borrower string
}

func (b *PhysicalBook) Booktype() BookType {
	return b.bookType
}

func (b *PhysicalBook) Name() string {
	return b.name
}

func (b *PhysicalBook) Author() string {
	return b.author
}

func (p *PhysicalBook) Borrow(borrower string) bool {
	// If there is no current borrower, allow the borrow and set borrower
	if p.borrower == "" {
		p.borrower = borrower
		return true
	} else {
		// Else do not allow borrow
		return false
	}
}

func (p *PhysicalBook) Return(borrower string) {
	if p.borrower != "" {
		p.borrower = ""
	}
}

func NewPhysicalBook(btype BookType, name string, author string) *PhysicalBook {
	return &PhysicalBook{btype, name, author, ""}
}

type DigitalBook struct {
	bookType  BookType
	name      string
	author    string
	limit     int
	borrowers []string
}

func (d *DigitalBook) Booktype() BookType {
	return d.bookType
}

func (d *DigitalBook) Name() string {
	return d.name
}

func (d *DigitalBook) Author() string {
	return d.author
}

func (d *DigitalBook) Borrow(borrower string) bool {
	// If borrow slot is available, append borrower to list
	if len(d.borrowers) < d.limit {

		d.borrowers = append(d.borrowers, borrower)
		return true

	} else {
		// Else do not allow borrow
		return false
	}
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func (d *DigitalBook) Return(borrower string) {
	index := 0
	for i := range d.borrowers {
		if d.borrowers[i] == borrower {
			index = i
			break
		}
	}

	RemoveIndex(d.borrowers, index)
}

func NewDigitalBook(btype BookType, name string, author string, limit int) *DigitalBook {
	return &DigitalBook{btype, name, author, limit, make([]string, 0)}
}

type Library struct {
	Books map[string]Book
	Users map[string]struct{}
}

func NewLibrary() *Library {
	return &Library{
		make(map[string]Book),
		make(map[string]struct{}),
	}
}

func (lib *Library) CheckUser(user string) bool {
	_, ok := lib.Users[user]
	return ok
}

func (lib *Library) AddUser(user string) {
	lib.Users[user] = struct{}{}
}

func (lib *Library) GetBook(bookName string) (Book, bool) {
	book, ok := lib.Books[bookName]
	return book, ok
}

func (lib *Library) AddBook(book Book) {
	lib.Books[book.Name()] = book
}

func (lib *Library) CheckBook(bookName string) Book {
	book := lib.Books[bookName]
	return book
}

func main() {
	var input int
	flag := true
	lib := NewLibrary()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("\t\t\t\t\t\t\tWELCOME TO CHETNA'S LIBRARY\n\n\n\n")

	for flag {
		fmt.Println("Choose from one of the options given:")
		fmt.Println("1. Add a NEW Book to the Library")
		fmt.Println("2. Add a NEW Member to our system")
		fmt.Println("3. Borrow a book")
		fmt.Println("4. Return a book")
		fmt.Println("5. Exit")
		fmt.Println()
		fmt.Printf("Enter your choice: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			var (
				choice int
				btype  BookType
			)

			fmt.Print("Enter the name of book: ")
			scanner.Scan()
			name := scanner.Text()
			fmt.Println()
			fmt.Println("Enter BookType: \n1. E-book\n2. AudioBook\n3. HardBack\n4. PaperBack\n5. Encyclopedia\n6. Magazine\n7. Comic: ")
			fmt.Scanln(&btype)
			fmt.Println()
			fmt.Print("Enter Author of book: ")
			scanner.Scan()
			author := scanner.Text()
			fmt.Println()
			fmt.Println("Enter choice of book: \n1. Physical\n2. Digital")
			fmt.Scanln(&choice)
			fmt.Println()

			if choice == 1 {
				lib.AddBook(NewPhysicalBook(btype, name, author))
			} else {
				var limit int
				fmt.Print("Enter the borrowing limit: ")
				fmt.Scanln(&limit)

				lib.AddBook(NewDigitalBook(btype, name, author, limit))
			}
			fmt.Println(lib)
			fmt.Println()

		case 2:
			var name string
			fmt.Print("Enter your name:")
			fmt.Scanln(&name)

			lib.AddUser(name)

			fmt.Println(lib)
			fmt.Println()

		case 3:
			var (
				username string
			)

			fmt.Print("Enter your name: ")
			fmt.Scanln(&username)

			ok := lib.CheckUser(username)
			if !ok {
				fmt.Println("User Does Not Exist")
				return
			} else {
				fmt.Println("You are registered! You may now proceed!")

				var bookname string
				fmt.Println("Enter the name of the book that you want to borrow: ")
				fmt.Scanln(&bookname)

				book, ok := lib.GetBook(bookname)
				if !ok {
					fmt.Println("Book Does Not Exist")
					return
				} else {
					if success := book.Borrow(username); success {
						fmt.Println("Book Borrowed Successfully!")
					} else {
						fmt.Println("Failure: Borrow Failed")
					}
				}
			}

		case 4:
			var (
				username string
			)

			fmt.Print("Enter your name: ")
			fmt.Scanln(&username)

			ok := lib.CheckUser(username)
			if !ok {
				fmt.Println("User Does Not Exist")
				return
			} else {
				fmt.Println("You are registered! You may now proceed!")

				var bookname string
				fmt.Print("Enter the name of the book that you want to return: ")
				fmt.Scanln(&bookname)

				book := lib.CheckBook(bookname)
				book.Return(username)

				fmt.Println("Book Returned Successfully")
			}

		default:
			flag = false
		}
	}
}
