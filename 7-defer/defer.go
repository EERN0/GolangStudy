// defer类似于析构函数，在函数体结束前触发

package main

import "fmt"

func main1() {
	// 写入defer关键字，在即将出函数体之前触发
	// 压栈的形式，先入后出
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")
}
