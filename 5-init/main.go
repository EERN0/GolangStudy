package main

/* 
import导包：
1. import _ "fmt"		给fmt包起一个别名，匿名，无法使用当前包的方法，但是会执行当前包内部的init()方法
2. import aa "fmt"		给fmt包起一个别名aa，aa.Println()来直接调用fmt包内的方法
3. import . "fmt"		(不推荐，多个包会出现函数名冲突)，将fmt包的方法全部导入到当前包中，直接用函数名就可以调用fmt包中的方法了，不用fmt.API来调用方法
*/

import (
	// 导包	gopath下路径要写完整
	_ "GolangStudy/5-init/lib1"		// 只会执行当前包内部的init()方法
	"GolangStudy/5-init/lib2"
)

func main() {
	// lib1.Lib1Test()
	lib2.Lib2Test()
	/*
		优先调用包的init()函数
		lib1. init() ...
		lib2. init() ...
		lib1Test()...
		lib2Test()...
	*/
}
