package main

import "fmt"

// 结构体字段可以通过结构体指针来访问，Go里的结构体指针，(*p).X和p.X等价

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9 // p是结构体指针，直接.结构体里面的字段
	fmt.Println(p.X)

	v2 := Vertex{Y: 10}       // Y显式赋值为10，X默认赋值为0
	fmt.Printf("v2:%v\n", v2) // v2:{0 10}
	p2 := &v2
	fmt.Println(p2) // &{0 10}
}
