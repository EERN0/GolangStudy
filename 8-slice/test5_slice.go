package main

import "fmt"

func main() {
	s := []int{1, 2, 3} // len=3, cap=3, [1,2,3]

	// 左闭右开区间，s1和s操作的是同一块内存，只是s1的可见范围比s小
	s1 := s[0:2]

	fmt.Println(s)  // [1 2 3]
	fmt.Println(s1) // [1 2]

	s1[0] = 100 // s1和s操作的是同一块内存

	fmt.Println(s)  // [100 2 3]
	fmt.Println(s1) // [100 2]

	s2 := make([]int, 4) // s2 = [0,0,0,0]
	// 将s中的值依次拷贝到s2中，s和s2操作的不是同一块内存
	copy(s2, s)
	fmt.Println(s2)		// s2 = [100 2 3 0]
}
