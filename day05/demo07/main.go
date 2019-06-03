package main

import (
	"fmt"
)

func main() {
	test()
}

func test() {
	//存放任意类型管道
	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "Oitm"
	cat := Cat{
		Name: "haha",
	}
	allChan <- cat

	//为了取第三个元素 需要将前两个元素先取出
	<-allChan
	<-allChan

	newCat := <-allChan
	fmt.Printf("newCat type = %T, value = %v \n", newCat, newCat)
	cat2 := newCat.(Cat)
	fmt.Println(cat2.Name)

}

type Cat struct {
	Name string
}

func testChan() {
	var intChan chan int

	//管道的数据长度不能超过初始化的容量
	intChan = make(chan int, 3)

	fmt.Printf("intChan的值=%v  intChan本身的地址=%p \n", intChan, &intChan)

	//向管道中写入数据
	intChan <- 10
	num := 211
	intChan <- num
	fmt.Printf("channel len=%v, cap=%v \n", len(intChan), cap(intChan))

	//从管道中取数据
	var num2 int = <-intChan //10 先进先出
	fmt.Println(num2)
}
