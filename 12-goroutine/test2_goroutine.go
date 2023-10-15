package main

import (
	"fmt"
	// "runtime"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("A.defer") // defer语句在return之前执行

		func() {
			defer fmt.Println("B.defer")
			// runtime.Goexit()		// 退出当前goroutine (内部方法退出，外部方法也跟着退出)

			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	go func(a int, b int) bool {
		fmt.Println("a = ", a, ", b = ", b)
		return true
	}(10, 20)

	// 死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
