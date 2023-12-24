package main

import "fmt"

// var a [10]int : 将变量a声明为有10个整数的数组
func foo1() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// 数组的长度是类型的一部分，有严格的数据类型，必须维度相同才能匹配
// 固定长度是值传递
func printArray1(myArray [5]int) {
	for idx, value := range myArray {
		fmt.Println("index = ", idx, ", value = ", value)
	}

	myArray[0] = 100
}

func main1() {
	// 固定长度的数组
	var myArray1 [10]int
	for i := 0; i < len(myArray1); i++ {
		fmt.Print(myArray1[i])
	}

	fmt.Println()

	// 遍历：基于范围的for循环
	myArray2 := [5]int{1, 2, 3, 4}
	for idx, value := range myArray2 {
		fmt.Println("index = ", idx, ", value = ", value)
	}

	fmt.Println("=================")
	// printArray1(myArray1)	// 报错，参数类型不匹配
	printArray1(myArray2)

	// 打印数组类型
	fmt.Printf("myArray1 types = %T\n", myArray1) // [10]int
	fmt.Printf("myArray2 types = %T\n", myArray2) // [5]int

}
