package main

import (
	"fmt"
	"runtime"
	"time"
)

// Go 自动提供了每个 case 后面所需的 break 语句， 除非以 fallthrough 语句结束，否则分支会自动终止。
// switch 的 case 无需为常量，且取值不必为整数

func foo1() {
	fmt.Println("Go run on")

	// runtime.GOOS获取当前操作系统
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}
}

// switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。
func foo2() {
	fmt.Println("When's Monday?")

	today := time.Now().Weekday()
	fmt.Printf("Type of today:%T, today is: %s\n", today, today.String())

	switch time.Monday {
	case today + 0:
		fmt.Println("Monday is today")
	case today + 1:
		fmt.Println("Monday is tomorrow")
	case today + 2:
		fmt.Println("Monday is in two days")
	default:
		fmt.Println("Monday is too far away")
	}
}

// 没有条件的switch语句比if-then-else语句更加清晰
func foo3() {
	t := time.Now()
	fmt.Println("now is", t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}

func main() {
	foo1()
	fmt.Println("----")

	foo2()
	fmt.Println("----")

	foo3()
}
