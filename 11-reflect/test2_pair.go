package main

import "fmt"

// 抽象接口，接口是指针
type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体的类
type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a Book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a Book")
}

func main2() {
	b := &Book{} // b: pair<type:Book, value:book{}地址>

	var r Reader // r: pair<type:, value:>
	r = b        // r: pair<type:Book, value:book{}地址>

	r.ReadBook()

	var w Writer   // w: pair<type:, value:>
	w = r.(Writer) // 断言转换成功，r原本是Reader接口类型转换成Writer接口类型。因为w、r具体的type是一致的

	w.WriteBook() // w: pair<type:Book, value:book{}地址>
}
