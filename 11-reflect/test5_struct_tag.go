package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name"  doc:"我的名字"`
	Sex  string `info:"sex"`
}

// 解析结构体类型中的标签
func findTag(str interface{}) {
	t := reflect.TypeOf(str) 			// Elem()得到当前结构体全部的元素，但是为啥要用Elem()? 传参是指针的时候，要加上.Elem()
	// t := reflect.TypeOf(str).Elem() 	// 实参是指针的时候，要加上.Elem()

	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")		// 
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", taginfo, ", doc: ", tagdoc)
	}
}

func main5() {
	var re resume
	findTag(re)		// 如果参数是&re，那么findTag()函数中必须是 t := reflect.TypeOf(str).Elem()
}
