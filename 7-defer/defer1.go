/*
	defer的执行顺序
*/
package main

import "fmt"

func fun1(){
	fmt.Println("A")
}

func fun2(){
	fmt.Println("B")
}

func fun3(){
	fmt.Println("C")
}

func main2(){
	// 入栈，先入后出	C	B	A
	defer fun1()
	defer fun2()
	defer fun3()
}