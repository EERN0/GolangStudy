package main

// import "fmt"
// import "time"
import (
	"fmt"
	"time"
)

func main() { // 函数的{一定是和函数名在同一行，否则编译错误
	fmt.Println("hello Go!")
	time.Sleep(1 * time.Second)

	fmt.Println(1 * time.Millisecond)
	fmt.Println(1 * time.Second)
}
