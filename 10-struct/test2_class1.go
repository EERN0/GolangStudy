package main

import "fmt"

// 类名首字母大写，表示其他包也能够访问
type Hero struct {
	// 类成员首字母大写，表示该成员是public；首字母小写则只能类内访问
	Name  string
	Ad    int
	Level int
}

// 类对象的成员方法
func (hero Hero) Show() {
	fmt.Println("Name = ", hero.Name)
	fmt.Println("Ad = ", hero.Ad)
	fmt.Println("Level = ", hero.Level)
}

// 传指针，操作的是同一个对象
func (hero *Hero) GetName() string {
	return hero.Name
}

func (hero *Hero) SetName(newName string) {
	hero.Name = newName
}

func main2() {
	// 创建一个对象
	hero := Hero{Name: "zhang3", Ad: 100, Level: 1}
	hero.Show()

	hero.SetName("li4")

	hero.Show()
}
