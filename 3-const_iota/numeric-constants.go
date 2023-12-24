package main

import "fmt"

const (
	Big   = 1 << 100  // int最大存储2^63-1
	Small = Big >> 99 // 2
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	var MaxInt int = 1<<63 - 1 // int: -2^63 ~ 2^63-1
	fmt.Println("MaxInt:", MaxInt)
	var MaxUint uint = 1<<64 - 1 // uint: 0 ~ 2^64-1
	fmt.Println("MaxUint:", MaxUint)

	fmt.Println(Small)        // 2
	fmt.Println(float64(Big)) // 1.2676506002282294e+30
}
