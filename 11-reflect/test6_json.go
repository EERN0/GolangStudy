package main

import (
	"encoding/json"
	"fmt"
)


// 结构体标签在json中的应用，定义标签为json，可以用json来序列化和反序列化
type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"xingye", "zhangbozhi"}}

	// 序列化，go类型序列化为字符串形式
	jsonStr, err := json.Marshal(movie) // jsonStr = {"title":"喜剧之王","year":2000,"rmb":10,"actors":["xingye","zhangbozhi"]}

	if err != nil {
		fmt.Println("json序列化失败: ", err)
	} else {
		fmt.Printf("jsonStr = %s\n", jsonStr)
	}

	// 反序列化，字符串反序列化为go格式
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie) // myMovie: {喜剧之王 2000 10 [xingye zhangbozhi]}

	if err != nil {
		fmt.Println("json反序列化失败: ", err)
	} else {
		fmt.Printf("myMovie: %v\n", myMovie)
	}
}
