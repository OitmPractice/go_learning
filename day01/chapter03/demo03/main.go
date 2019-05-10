package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var a = false
	fmt.Println(a)
	fmt.Printf("a = %d\n", unsafe.Sizeof(a)) // a =1 , bool 类型占用1个字节

	//反引号输出复杂字符串
	str := `
		import (
			"fmt"
			"unsafe"
		)
	`
	fmt.Println(str)

	//字符串拼接方式
	str2 := "Hello" + " World" +
		" Hello" + " World"
	str2 += " Oitm"
	fmt.Println(str2)

	//基本数据类型的默认值
	//整型 0
	//浮点型 0
	//字符串 ""
	//布尔类型 false

	var b int
	var c float32
	var d float64
	var e bool
	var f string
	// %v 按照变量的原值输出
	fmt.Printf("b=%d  c = %f  d = %f  e = %v  f = %v", b, c, d, e, f)
	//b=0  c = 0.000000  d = 0.000000  e = false  f =
}
