package model

type Person struct {
	Name string
	Age  int
}

type Human struct {
	Name   string
	Height float32
}

type People struct {
	Weight float32
	Name   string
}

type Total struct {
	Person
	Human
	P People //有名继承时，需要对象.p.name   不能省略p
	A int
}
