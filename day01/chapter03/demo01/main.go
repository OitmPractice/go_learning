package main

import (
	"fmt"
)

var (
	num1 = 100
	num2 = 200
	name = "Oitm_"
)

func main() {

	fmt.Println(num1, num2, name)

	var i int
	i = 10
	fmt.Println("i=", i)

	var str string
	str = "Oitm"
	fmt.Println(str)

	name := "Oitm be rich"
	fmt.Println(name)

	append := name + str
	fmt.Println(append)

}
