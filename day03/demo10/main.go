package main

import (
	"fmt"
)

//如果结构体的字段类型是：指针、slice、map时，他们的零值都是nil，如果需要使用这样的字段，需要先make分配内存空间
func main() {
	students := make(map[string]Student, 10)
	stu1 := Student{"Oitm", 18, "上海"}
	students["key"] = stu1
	fmt.Println(students)

	var stu2 Student
	stu2.Name = "Oitm2"
	stu2.Age = 18
	stu2.Address = "魔都"
	fmt.Println(stu2)

	test(stu2)
	fmt.Println(stu2)

	var students2 []Student
	students2 = append(students2, stu1)
	fmt.Println("students2 ==  ", students2)
}

//struct值类型测试
func test(stu Student) {
	stu.Name = "momo"
	fmt.Println(stu)
}

//定义一个学生结构体
type Student struct {
	Name    string
	Age     int
	Address string
}
