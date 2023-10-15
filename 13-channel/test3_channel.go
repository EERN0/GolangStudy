package main

import "fmt"

// 关闭channel
func main3() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		// close可以关闭channel
		close(c)
	}()

	/* 	for {
		// ok为true表示channel没有关闭，如果为false表示channel已经关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	} */

	// 可以使用range来迭代操作channel
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Main Finshed...")
}
