package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	//结构体字段上添加标签
	monster := Monster{"牛魔王", 168, "芭蕉扇"}

	//序列化结构体为json字符串
	jsonByte, _ := json.Marshal(monster)
	fmt.Println(string(jsonByte))

}
