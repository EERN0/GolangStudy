package main

import (
	"fmt"
	"time"
)

func main2() {
	c := make(chan int, 3) // 有缓冲的channel。当channel已经满了，再向里面写数据，就会阻塞；当channel为空，从里面取数据也会阻塞

	fmt.Println("len(c):", len(c), ", cap(c):", cap(c)) // len(c): 0 , cap(c): 3

	go func() {
		defer fmt.Println("子go程结束")

		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("子go程正在运行, 发送的元素=", i, "len(c)=", len(c), ", cap(c)=", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	// 从管道中取出元素
	for i := 0; i < 3; i++ {
		num := <-c // 从c中接收数据，并赋值给num
		fmt.Println("num =", num)
	}

	fmt.Println("main 结束")
}
