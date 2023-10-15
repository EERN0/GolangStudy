package main

import "fmt"

// 变长数组，引用传递
func printArray(myArray []int) {
	// _表示匿名的变量
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
	myArray[0] = 100
}

func main2() {
	// 动态数组，切片 slice
	myArray := []int{1, 2, 3, 4}

	fmt.Printf("myArray type is %T\n", myArray) // []int

	printArray(myArray)

	fmt.Println("==============")

	printArray(myArray) // myArray[0]变为100
}
