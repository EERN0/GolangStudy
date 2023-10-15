package main

import (
	"fmt"
	"time"
)

// 子goroutine执行
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}


// 主goroutine
func main1() {
	// 创建子go携程	去执行newTask()
	go newTask()

	i := 0

	for {
		i++
		fmt.Printf("main goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
