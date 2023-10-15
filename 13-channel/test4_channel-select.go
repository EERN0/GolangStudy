/* 
	select能监控多路channel的状态
*/
package main

import "fmt"

func fib(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x: // 如果channel-c可写，就会执行这个case
			x, y = y, x+y
		case <-quit: // 如果quit有数据，就会执行
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	// 子go程
	go func() {
		for i := 0; i < 6; i++ {
			// 取出channel-c中的数据，并打印
			fmt.Println(<-c)
		}

		// 取完6次数据后，往channel-quit中写入0，当slelect发现quit中有数据时，说明要结束了
		quit <- 0
	}()

	// main go程
	fib(c, quit)
}
