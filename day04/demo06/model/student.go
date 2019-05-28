package model

import (
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Age   int
	score int
}

func (stu *Student) ShowInfo() {
	fmt.Printf("name = [%v], age = [%v], score = [%v] \n", stu.Name, stu.Age, stu.score)
}

func (stu *Student) SetScore(score int) {
	(*stu).score = score
}

type Pupil struct {
	Student
	Height float64
}

func (p *Pupil) testing() {
	fmt.Println("小学生在考试")
}

type Graduate struct {
	Student
	Pupil
}

func (g *Graduate) testing() {
	fmt.Println("大学生在考试")
}
