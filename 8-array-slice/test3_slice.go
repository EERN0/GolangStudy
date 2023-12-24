package main

import "fmt"

// slice声明的几种方式
func main3() {
	// 1.声明一个空切片slice1，没有分配空间
	// var slice1 []int
	// fmt.Printf("len = %d", len(slice1), ", %v", slice1)		// slice1未分配空间

	// 2.切片slice2，初始化，默认值为1 2 3，长度len是3
	slice2 := []int{1, 2, 3}
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)

	// 3.给切片slice3分配3个空间，默认初始值为0
	slice3 := make([]int, 3)
	fmt.Printf("len = %d, slice = %v\n", len(slice3), slice3)

	var slice4 []int = make([]int, 4)
	fmt.Printf("len = %d, slice = %v\n", len(slice4), slice4)

	var slice1 []int		// slice1是空切片
	// 判断slice是否为空
	if slice1 != nil {
		fmt.Println("slice1不是一个空切片")
	} else {
		fmt.Println("slice1是一个空切片")
	}
}
