package main

import (
	"fmt"
)

func main1() { 
	// 1.第一种方式，声明一个空map，使用前必须要用make来分配空间
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 是一个空map")
	}

	// 在使用map前，需要先用make给map分配空间，否则报错
	myMap1 = make(map[string]string, 5)
	myMap1["a"] = "java"
	myMap1["b"] = "c++"
	myMap1["c"] = "go"
	fmt.Println(myMap1) // map[a:java b:c++ c:go]

	// 2.第二种方式，直接用make来声明并初始化map
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "go"

	fmt.Println(myMap2) // map[1:java 2:c++ 3:go]

	// 3.第三种方式, 在初始化的时候已经开辟空间了，不再需要make
	myMap3 := map[string]string{
		"a": "java",
		"b": "c++",
		"c": "go",
	}
	fmt.Println(myMap3)
}
