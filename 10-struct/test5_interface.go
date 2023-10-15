package main

import "fmt"

// interface{}是万能数据类型，可以接收任何类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	// interface{}万能类型如何区分，表示的数据类型是什么？
	// interface{} 提供 "类型断言" 机制
	value, ok := arg.(string) // arg.(string)：判断arg是不是string类型

	if !ok {
		fmt.Println(arg, "is not string type")
		fmt.Printf("%v type is %T\n", arg, arg)
	} else {
		fmt.Println("arg is string type, value = ", value)
	}
}

type MyBook struct {
	title string
}

func main() {
	book := MyBook{"Golang"}
	myFunc(book)
	fmt.Println("----------------")

	myFunc(3.14)
	fmt.Println("----------------")

	myFunc("aaa")
}
