package main

import (
	"fmt"
)

type Student struct {
	name string
	age  int
}

func (stu Student) say() string {
	return fmt.Sprintf("name = [%v], age = [%v]", stu.name, stu.age)
}

func main() {
	stu := Student{
		name: "Oitm",
		age:  18,
	}
	fmt.Println(stu.say())

}
