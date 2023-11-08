package main

import (
	"fmt"
)

// 打印斐波那契数列
func fib(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x: // channel-c可写（c中数据被取出来了，没有数据就可写），就执行这个case，往c中写入数据
			x, y = y, x+y
		case <-quit: // channel-quit有数据的话，直接退出
			fmt.Println("quit")
			return
		}
	}
}

func main1() {
	c := make(chan int)		// 未指定channel缓冲区大小，不能直接往里面写数据	直接运行 [c<-10] 会报错
	quit := make(chan int)

	// 子go程序 goroutine负责取出channel-c中的数据
	go func() {
		for i := 0; i < 10; i++ { // c中有数据，就取出channel-c中的数据，执行10轮
			fmt.Print(<-c, " ") // channel中没有数据，就会阻塞，等有数据后，才能把数据取出来
		}
		// 打印10轮结束后，往quit中写数据
		// select监测到quit中有数据后，直接退出
		quit <- 0
	}()

	// 主go程
	fib(c, quit)
}
