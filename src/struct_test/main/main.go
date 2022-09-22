package main

import "fmt"

// 使用type声明类型

type myint int // 声明一个myint类型, 这个类型是int的别名

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
