package main

import "fmt"

// go语言中没有while语句
// for 初始化语句; 条件语句; 后置语句 {...}		初始化和后置语句都可以省略，只留下条件语句相当于C++中的while

func main1() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum) //5050

	// for中的初始化语句和后置语句可以省略，当作while使用
	for sum < 5050+1 {
		sum += sum
	}
	fmt.Println(sum) //10100

	// // 死循环，省略判断条件
	// for {
	// 	//...
	// }

	// 基于范围的遍历
	slice1 := []int{1, 2, 3}
	for idx := range slice1 {
		fmt.Println("index:", idx, "value:", slice1[idx])
	}
}
