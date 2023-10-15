package main

import "fmt"

// 数组有严格的数据类型，必须维度相同才能匹配
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
	fmt.Printf("myArray1 types = %T\n", myArray1) 	// [10]int
	fmt.Printf("myArray2 types = %T\n", myArray2) 	// [5]int

}
