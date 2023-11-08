package main

import (
	"fmt"
	"time"
)

func main() {
	isAlive := make(chan bool, 1) // 指定channel缓冲区大小为1，才能写入数据。否则报错
	// isAlive <- true

	fmt.Println("hello")
	fmt.Println(isAlive)

	select { // IO多路复用，同时监控多个channel，有一个case满足就执行那个case，不然一直阻塞
	case <-isAlive:
		fmt.Println("isalive: true")
		close(isAlive)

	case <-time.After(time.Second * 5):
		fmt.Println("阻塞了5s...")
	}
}
