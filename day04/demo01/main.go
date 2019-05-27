package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

type Rect struct {
	leftUp, rightDown Point
}

type A struct {
	Num int
}

type B struct {
	Num int
}

func main() {
	//结构体中属性的内存分配是连续的
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	fmt.Println(r1)

	//结构体强转  两者的属性必须一致
	b := B{20}
	c := A(b)
	fmt.Println(c)
}
