package main

import "fmt"

func main1() {
	var a string
	a = "123qwe" // pair<type:string, value:"123qwe">

	var allType interface{} // 万能类型
	allType = a             // pair<type:string, value:"123qwe">

	str, _ := allType.(string) // 断言，检查allType是否是string类型
	fmt.Println(str)
}
