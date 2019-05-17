package main

import (
	"errors"
	"fmt"
)

func main() {
	test()
	fmt.Println("Oitm")

	test02()

}

func test02() {
	err := customErr("Oitm")
	if err != nil {
		panic(err) //如果发生错误，就输出，并终止程序
	}
	fmt.Println("test02()到这里了")

}

//自定义异常
func customErr(name string) error {
	if name == "Oitm" {
		return nil
	}
	return errors.New("name error！")
}

func test() {
	// 错误处理机制
	// 使用defer+recover来捕获和处理异常
	defer func() {
		//recover()内置函数，可以捕获到异常
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	temp := 0
	num1 := 10 / temp
	fmt.Println(num1)
}
