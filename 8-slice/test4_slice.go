package main

import "fmt"

func main4() {
	var numbers = make([]int, 3, 4)                                                     // 长度为3，容量为4
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers) // len = 3, cap = 4, slice = [0 0 0]

	// 切片尾部追加一个元素10
	numbers = append(numbers, 10)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers) // len = 4, cap = 4, slice = [0 0 0 10]

	// len==capacity后，再追加元素，引起扩容，扩容2倍
	numbers = append(numbers, 20)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers) // len = 5, cap = 8, slice = [0 0 0 10 20]

	fmt.Println("==========")

	var numbers2 = make([]int, 3)	// len和cap大小都是3
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2) // len = 3, cap = 3, slice = [0 0 0]
	numbers2 = append(numbers2, 100)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2) // len = 4, cap = 6, slice = [0 0 0 100]
}
