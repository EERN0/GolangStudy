package main

import "fmt"

// 切片的长度 len() 是它所包含元素的个数
// 切片的容量 cap() 是从它的第一个元素开始（切片指针），到其底层数组元素末尾的个数
func foo1() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为0
	s = s[:0]
	printSlice(s) // len=0 cap=6 slice=[]

	// 扩展切片长度
	s = s[:4]
	printSlice(s) // len=4 cap=6 slice=[2 3 5 7]

	// 丢弃前两个元素
	s = s[2:]
	printSlice(s) // len=2 cap=4 slice=[5 7]
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(s), cap(s), s)
}

func main() {
	foo1()
	fmt.Println("--------")
	// slice扩容cap变为两倍
	var numbers = make([]int, 3, 4)                                                     // 长度为3，容量为4
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers) // len = 3, cap = 4, slice = [0 0 0]

	// 切片尾部追加一个元素10
	numbers = append(numbers, 10)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers) // len = 4, cap = 4, slice = [0 0 0 10]

	// len==capacity后，再追加元素，引起扩容，扩容2倍
	numbers = append(numbers, 20)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers) // len = 5, cap = 8, slice = [0 0 0 10 20]

	fmt.Println("==========")

	var numbers2 = make([]int, 3)                                                          // len和cap大小都是3
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2) // len = 3, cap = 3, slice = [0 0 0]
	numbers2 = append(numbers2, 100)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2) // len = 4, cap = 6, slice = [0 0 0 100]
}
