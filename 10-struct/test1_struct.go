package main

import "fmt"

type Book struct {
	title  string
	author string
}

// 结构体作为参数，值传递
func changeBook(book Book) {
	book.author = "66"
}

// 结构体传指针
func changeBook2(book *Book) {
	book.author = "555"
}

func main1() {
	var book1 Book
	book1.title = "golang"
	book1.author = "zhang3"

	fmt.Printf("%v\n", book1)		// {golang zhang3}
	
	changeBook(book1)	// 值传递，只改变副本

	fmt.Printf("%v\n", book1)		// {golang zhang3}
	
	changeBook2(&book1)	// 传指针

	fmt.Printf("%v\n", book1)		// {golang 555}
}
