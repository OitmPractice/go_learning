package main

import (
	"fmt"
	"go_code/day04/demo06/model"
)

func main() {
	pupil := &model.Pupil{}
	pupil.Student.Name = "Oitm"
	pupil.Student.Age = 18
	pupil.SetScore(100)

	pupil.ShowInfo()
	pupil.Name = "Oitm2"
	pupil.ShowInfo()
	/*
		继承注意
		结构体在本包内 可以使用嵌套匿名结构体的所有字段和方法
	*/

	g := model.Graduate{}
	g.Height = 178
	g.ShowInfo()
	g.Age = 18
	g.Name = "Oitm"
	g.Student.Name = "Oitm1"
	g.Pupil.Name = "Oitm2"
	fmt.Println(g.Height)
	fmt.Println(g.Name)
	fmt.Println(g.Student.Name)
	fmt.Println(g.Pupil.Name)

	person := model.Total{}
	person.Person.Name = "oitm"

	person.P.Weight = 10
	person.A = 100

	fmt.Println(person)

}
