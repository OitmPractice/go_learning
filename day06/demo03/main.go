package main

import (
	"fmt"
	"reflect"
)

func main() {
	monster := Monster{
		Name:  "猪八戒",
		Age:   100,
		Sex:   "男",
		Score: 59,
	}
	TestStruct(monster)
}

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("-------start")
	fmt.Println(s)
	fmt.Println("-------end")
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {

	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	kd := v.Kind()
	if kd != reflect.Struct {
		fmt.Println("传入对象不是结构体")
		return
	}

	//遍历所有字段 及其值
	fieldNum := v.NumField()
	for i := 0; i < fieldNum; i++ {
		tag := t.Field(i).Tag.Get("json")
		if tag != "" {
			fmt.Println("tag == ", tag)
		}
		fmt.Println("val == ", v.Field(i))
	}

	//获取结构体的 字段
	tFieldNum := t.NumField()
	for i := 0; i < tFieldNum; i++ {
		fmt.Println("tField = ", t.Field(i))
	}

	//获取结构体对象有多少方法
	vMethodNum := v.NumMethod()
	for i := 0; i < vMethodNum; i++ {
		fmt.Println("vMethod == ", v.Method(i).String())
	}

	//获取结构体的所有方法 这个能获取到方法的方法名字
	tMenthodNum := t.NumMethod()
	for i := 0; i < tMenthodNum; i++ {
		fmt.Println("tMethod == ", t.Method(i).Name)
	}

	//动态调用方法
	//Method 中方法的排序方式根据方法的首字母 按照ASCII码排序
	v.Method(1).Call(nil) //所以这里是Print

	//动态调用带参数的方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	res := v.Method(0).Call(params)
	fmt.Println(res[0].Int())

}
