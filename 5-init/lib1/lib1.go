package lib1

import "fmt"

// 包对外提供的函数的函数名要大写，权限为public，可以在别的包里面访问
func Lib1Test() {
	fmt.Println("lib1Test()...")
}

func init() {
	fmt.Println("lib1. init() ...")
}
