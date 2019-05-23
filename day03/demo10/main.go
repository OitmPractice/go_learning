package main

import (
	"fmt"
)

func main() {
	students := make(map[string]Student, 10)
	stu1 := Student{"Oitm", 18, "上海"}
	students["key"] = stu1
	fmt.Println(students)
}

//定义一个学生结构体
type Student struct {
	Name    string
	Age     int
	Address string
}
