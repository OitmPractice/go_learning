package main

import (
	"fmt"
	"sort"
)

type Usb interface {
	//声明两个没有实现的方法
	Start()
	Stop()
}

type phone []Phone

type Phone struct {
	Name   string
	Number int
}

func (p phone) Len() int {
	return len(p)
}

//Less方法就是决定你使用什么标准进行排序
//1. 按number的年龄从小到大排序!!
func (p phone) Less(i, j int) bool {
	return p[i].Number > p[j].Number
}

// Swap swaps the elements with indexes i and j.
func (p phone) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	//做一个自定义结构体排序功能
	var phones phone
	phones = append(phones, Phone{
		Name:   "Oitm",
		Number: 10010,
	})

	phones = append(phones, Phone{
		Name:   "Momo",
		Number: 10086,
	})

	sort.Sort(phones)
	fmt.Println(phones)
}
