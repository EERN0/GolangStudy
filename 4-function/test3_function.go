package main

import "fmt"

// 返回值：一个int
func foo1(a string, b int) int {
	fmt.Println("---foo1---")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100
	return c
}

// 多个返回值，匿名的
func foo2(a string, b int) (int, int) {
	fmt.Println("---foo2---")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 666, 777
}

// 多个返回值，有形参名的
func foo3(a string, b int) (r1, r2 int) {
	fmt.Println("---foo3---")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// r1 r2是形参，初始化默认值是0
	fmt.Println("r1 = ", r1, ", r2 = ", r2)

	r1 = 100
	r2 = 200
	return
	// 直接[return 100, 200]都行
}

func main() {
	c := foo1("abc", 555)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("abc", 111)
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2) // ret1 =  666 ret2 =  777

	ret3, ret4 := foo3("xxx", 222)
	fmt.Println("ret3 = ", ret3, "ret4 = ", ret4)

}
