package main

import "fmt"


// Book 定义结构体
type Book struct {
	title string
	auth  string
}

func changeBook(book *Book) {
	book.auth = "777"
}

func main() {
	var book1 Book
	book1.title = "golang"
	book1.auth = "soulchild"

	fmt.Println(book1)
	changeBook(&book1)
	fmt.Println(book1)

}
