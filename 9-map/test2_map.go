package main

import "fmt"

// map作为参数是引用传递
func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("key =",key, ", value =", value)
	}
}

func main() {
	cityMap := make(map[string]string)

	// 添加
	cityMap["china"] = "beijing"
	cityMap["japan"] = "tokyo"
	cityMap["usa"] = "newyork"

	// 删除
	delete(cityMap, "china") // 直接删除这一行的

	// 修改
	cityMap["usa"] = "aaaa"

	printMap(cityMap)

	fmt.Println("-------------")

	
}
