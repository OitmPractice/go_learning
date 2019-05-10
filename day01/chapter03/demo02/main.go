package main

import (
	"fmt"
	"unsafe"
)

var (
	num1 = 100
	num2 = 200
	name = "Oitm_"
)

func main() {
	//Printf 格式化输出
	//%T 查看数据类型
	fmt.Printf("num1 数据类型是 %T \n", num1)

	//查看变量占用字节数
	fmt.Printf("name 数据类型是 %T \n", name)
	fmt.Printf("name 占用的字节数 %d \n", unsafe.Sizeof(name))

	var a byte = 'a'
	var b byte = '0'
	fmt.Println("a = ", a, " b =", b)
	fmt.Printf("a = %c\n", a)

	var c int = '冉'
	fmt.Printf("c = %c \n", c)//c = 冉

	var d int = 97
	fmt.Printf("d = %c \n", d) // 输出a

	// 字符类型可进行运算  相当于一个整数，运算时按照码值进行运算
	var e = 10 + 'a'
	fmt.Println("e = ", e) // e = 107
}
