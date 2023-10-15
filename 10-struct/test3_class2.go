package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (human *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (human *Human) Walk() {
	fmt.Println("Huaman.Walk()...")
}

/* ============================================== */

type SuperMan struct {
	Human // SuperMan继承了Human类的方法

	level int
}

// 子类重写父类Eat()方法
func (superman *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

// 子类的新方法
func (superman *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (superman *SuperMan) Print() {
	fmt.Println("name = ", superman.name)
	fmt.Println("sex = ", superman.sex)
	fmt.Println("level = ", superman.level)
}

func main3() {
	h := Human{"zhang3", "man"}
	h.Eat()
	h.Walk()

	// 定义一个子类对象
	// 1.方法一：先构造一个父类对象，再添加子类属性
	// s := SuperMan{Human{"li4", "female"}, 100}

	// 2.方法二：
	var s SuperMan
	s.name = "li4"
	s.sex = "female"
	s.level = 100

	s.Walk() // 调用的是父类的方法
	s.Eat()  // 子类重写的方法
	s.Fly()  // 子类的方法

	s.Print()
}
