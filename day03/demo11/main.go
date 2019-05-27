package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	//方式1
	var p1 Person
	p1.Name = "Oitm"

	//2
	p2 := Person{"song", 18}
	p2.Age = 18

	//3
	var p3 *Person = new(Person)
	(*p3).Name = "momo"
	//Go设计者为了程序员方便，底层会对 p3.Age = 18进行处理，加上取值运算（*p3）
	p3.Age = 18

	//4
	var p4 *Person = &Person{}
	(*p4).Name = "shu"
	p4.Age = 18
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(*p3)
	fmt.Println(*p4)

}
