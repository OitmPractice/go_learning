package main

type Usb interface {
	//声明两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
}

func main() {
	//做一个自定义结构体排序功能
}
