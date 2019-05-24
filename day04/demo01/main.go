package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}

	fmt.Println(&r1.leftUp.x)
	fmt.Println(&r1.leftUp.y)
	fmt.Println(&r1.rightDown.x)
	fmt.Println(&r1.rightDown.y)
}
