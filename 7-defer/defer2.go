/*
	知识点2：defer语句在return语句之后执行!!
*/

package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called...")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	// return func called...
	// defer func called...
	returnAndDefer()
}
