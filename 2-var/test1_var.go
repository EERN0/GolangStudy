package main

/*
	四种变量的声明方式
*/

import (
	"fmt"
)

// 【声明全局变量】	方法一、方法二、方法三是可以的
var gA float64 = 20.4
var gB = "xxx"

// 方法四:=不能用来声明全局变量，只能用在函数体内。在函数体外，只能用func、var等关键字开头。
// gC := 10 	// := 声明全局变量直接报错

func main() {
	// 方法一：声明一个变量	默认值是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	// 方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	// 方法三：初始化时可以省略数据类型，可以自动推导
	var c = "abcd"
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c) // type of c = string

	// 方法四：省略var关键字，直接用 := 来自动匹配
	// := 只能用在 函数体内来声明
	d := 10.11
	fmt.Println("d = ", d)
	fmt.Printf("type of d = %T\n", d)

	fmt.Println("gA = ", gA, ", gB = ", gB)

	// ===================================
	// 声明多个变量
	var xx, yy int = 1, 2
	fmt.Println("xx = ", xx, ", yy = ", yy)

	// 多变量不同类型，自动推导
	var kk, ll = 100, "abcdef"
	fmt.Println("kk = ", kk, ", ll = ", ll)

	// 使用:=自动推导
	mm, nn := 20, "hello"
	fmt.Println("mm = ", mm, ", nn = ", nn)

	// 多行的多变量声明
	var (
		vv      = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, ", jj = ", jj)
}
