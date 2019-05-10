package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 100
	var n1 float32 = float32(i)
	fmt.Printf("i = %v   n1=%v", i, n1)

	//基本数据类型和string之间的转换
	var num1 int = 100
	var num2 float64 = 12.34
	var b bool = true
	var myChar byte = 'O'

	var str string
	str = fmt.Sprintf("%d", num1)
	fmt.Println("str = ", str)

	str = fmt.Sprintf("%f", num2)
	fmt.Println("str = ", str)

	str = fmt.Sprintf("%t", b)
	fmt.Println("str = ", str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Println("str = ", str)

	//第二种
	str = strconv.FormatInt(100, 10)
	fmt.Println("str = ", str)

	///string 转基本类型
	str = "true"
	boo, _ := strconv.ParseBool(str)
	fmt.Println(boo)

	
}
