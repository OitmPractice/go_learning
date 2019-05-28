package main

import (
	"fmt"
)

type A struct {
	Num int
}

//在方法名前指定调用的方法类型 和函数的区别
func (a A) test() {
	fmt.Println(a.Num)
}

type Person struct {
	Name string
	Age  int
}

//值传递
func (p Person) speak() {
	p.Age += 1
	fmt.Println(p.Age)
	fmt.Println(p.Name + " is a good man")
}

//指针copy
func (p *Person) cook() {
	fmt.Println(p.Age)
	p.Age += 10
	fmt.Println(p.Name + " can cook")
}

func (p Person) sum(a int, b int) int {
	return a + b
}

func main() {
	a := A{10}
	a.test()

	p := Person{"Oitm", 18}
	p.speak()
	fmt.Println(p.Age)

	sum := p.sum(11, 12)
	fmt.Printf("sum == %d \n", sum)

	p.cook()
	fmt.Println(p.Age)
}
