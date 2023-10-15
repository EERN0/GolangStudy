/* 
	多态
*/

package main

import (
	"fmt"
)

// 接口，是一个指针
// 相当于一个父类指针，指向子类对象，实现多态
type Animal interface {
	Sleep()
	GetColor() string // 获取动物的颜色
	GetType() string  // 获取动物的种类
}

// 在具体的类中重写接口的所有方法，才算实现了这个接口

// cat类
type Cat struct {
	color string
}

func (cat *Cat) Sleep() {
	fmt.Println("Cat is sleeping...")
}

func (cat *Cat) GetColor() string {
	return cat.color
}

func (cat *Cat) GetType() string {
	return "cat"
}

// dog类
type Dog struct {
	color string
}

func (dog *Dog) Sleep() {
	fmt.Println("Dog is sleeping...")
}

func (dog *Dog) GetColor() string {
	return dog.color
}

func (dog *Dog) GetType() string {
	return "dog"
}

// 多态，父类指针指向子类对象
func showAnimal(animal Animal) {
	animal.Sleep()
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("type = ", animal.GetType())
}

func main4() {
	// 接口类型是指针类型，用父类指针指向子类对象，实现多态
	var animal Animal
	animal = &Cat{"green"}
	animal.Sleep()
	fmt.Println(animal.GetColor())
	fmt.Println(animal.GetType())

	fmt.Println("---------")

	animal2 := Dog{"yellow"}
	animal2.Sleep()
	fmt.Println(animal2.GetColor())
	fmt.Println(animal2.GetType())

	fmt.Println("---------")

	showAnimal(&Dog{"black"})
}
