package main

import (
	"fmt"
	"go_code/day04/demo05/model"
)

func main() {
	var stu = model.NewStudent("oitm", 100)
	fmt.Println(*stu)

	fmt.Println(stu.GetScore())
	stu.SetScore(123)
	fmt.Println(stu.GetScore())

}
