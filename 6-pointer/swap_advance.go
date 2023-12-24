package main

import "fmt"

// 交换x,y两个变量，不使用中间变量

// 方法一：减减再加
func swap1(x, y int) (int, int) {
	x = y - x
	y = y - x
	x = y + x
	return x, y
}

// 方法二：异或 ⊕，二进制运算，相同为0，相异为1
func swap2(x, y int) (int, int) {
	x = x ^ y
	y = x ^ y // x^y^y = x^0 = x(1异或0为1，0异或0为0)
	x = x ^ y
	return x, y
}

// 方法三：使用栈，先进后出。但是Go里面得自己实现栈

func main() {
	x, y := swap1(10, 20)
	fmt.Println("x:", x, "y:", y) // x: 20 y: 10

	m, n := swap2(10, 20)
	fmt.Println("m:", m, "n:", n) // m: 20 n: 10

	// Go其实可以直接交换两变量的值
	x, y = y, x
	fmt.Println("x:", x, "y:", y) // x: 10 y: 20
}
