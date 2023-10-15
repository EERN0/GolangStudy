// 常量
package main

import (
	"fmt"
)

// const来定义枚举类型
// iota只能出现在const()中，只有在const有累加效果，iota【每一行】累加1
const (
	// 可以在const中添加一个关键字iota，[每行]iota累加1，第一行的iota的默认值是0
	BeiJing  = 10 * iota // iota = 0, BeiJing = 10*0
	ShangHai             // iota = 1, ShangHai = 10*1
	ShenZhen             // iota = 2, ShenZhen = 10*2
)

const (
	a, b = iota + 1, iota + 2 // iota=0, a=iota+1=1, b=iota+2=2
	c, d                      // iota=1, c=iota+1=2, d=iota+2=3
	e, f                      // iota=2, e=iota+1=3, d=iota+2=4

	g, h = iota * 2, iota * 3 // iota=3, g=iota*2=6, h=iota*3=9
	i, k                      // iota=4, i=iota*2=8, k=iota*3=12
)

func main() {
	const length = 10

	fmt.Println("length = ", length)
	// length = 100		// 常量不允许修改

	fmt.Println("====================")

	fmt.Println("BeiJing = ", BeiJing)
	fmt.Println("ShangHai = ", ShangHai)
	fmt.Println("ShenZhen = ", ShenZhen)

	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)
	fmt.Println("e = ", e, "f = ", f)
	fmt.Println("g = ", g, "h = ", h)
	fmt.Println("i = ", i, "k = ", k)

	// iota只能配合const()一起使用，iota只有在const有累加效果
	// var a int = iota			// 报错
	// fmt.Println(a)
}
