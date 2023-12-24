package main

import (
	"fmt"
	"math"
)

// if不用小括号，但是要{}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// if可以在判断语句前，执行一个简单语句，该语句声明的变量作用域仅在if、else之内
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v // v的作用域只在if的{}里面
	} else if v == lim {
		fmt.Printf("%g = %g\n", v, lim)
	} else {
		fmt.Printf("%g > %g\n", v, lim)
	}
	// 这里开始就不能使用局部变量v了
	return lim
}

func main() {
	// fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 9),  // 9
		pow(3, 2, 10), // 9
		pow(3, 3, 10), // 10
	)

	x, y := 10, 20
	x, y = y, x
	fmt.Println("x=", x, "y=", y)
}
