package main

import "fmt"

// 无缓冲channel
func main1() {
	// 定义一个channel
	c := make(chan int)		// 无缓冲的channel，双方都存在阻塞

	go func() {
		defer fmt.Println("goroutine结束")

		fmt.Println("goroutine 正在运行...")

		c <- 666 // 将 666 发送给c
	}()

	num := <-c 	// 从c中接受数据，并赋值给num
	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束...")
}
