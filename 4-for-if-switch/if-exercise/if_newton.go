package main
// 牛顿法开根号
import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	cnt := 1
	for {
		temp := z
		z -= (z*z - x) / (2 * z)

		if math.Abs(temp-z) < 1e-7 {
			break
		}
		fmt.Printf("第%d次迭代, x=%g, z=%g\n", cnt, x, z)
		fmt.Println("----")
		cnt++
	}
	return z
}

func main() {
	fmt.Println(
		Sqrt(2),
		Sqrt(9),
		Sqrt(78),
		Sqrt(100),
	)
}
