package main

// 类型 []T 表示一个元素类型为 T 的切片。
// 切片就像数组的引用，它不存储任何数据，只是描述了底层数组中的一段。
// 更改切片的元素会修改其底层数组中对应的元素，与它共享底层数组的切片都会被修改

import "fmt"

// 变长数组，引用传递
func printArray(myArray []int) {
	// _表示匿名的变量
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
	myArray[0] = 100
}

// 1.slice就是数组的引用，修改slice会修改底层数组，与该数组相关的所有切片都能观测到修改
func foo1() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	// 切片就是数组的引用
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) // [John Paul] [Paul George]

	// 修改slice，会修改到底层数组，该数组的所有切片都会变
	b[0] = "XX"
	fmt.Println(a, b)  // [John XX] [XX George]
	fmt.Println(names) // [John XX George Ringo]
}

// 2.创建切片时，会先创建一个数组，然后构建一个引用该数组的切片
func foo2() {
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
	}
	fmt.Println(s) // [{2 true} {3 false} {5 true} {7 true} {11 false}]
}

func main() {
	// 先会创建一个数组 [4]int{1,2,3,4}，然后构建一个引用了数组的切片
	myArray := []int{1, 2, 3, 4}
	fmt.Printf("myArray type is %T\n", myArray) // []int

	printArray(myArray)
	fmt.Println("==============")

	printArray(myArray) // myArray[0]变为100
	fmt.Println("==============")

	foo1()

	foo2()
}
