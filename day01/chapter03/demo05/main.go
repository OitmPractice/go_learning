package main

import (
	"fmt"
)

func main() {
	num := 100
	fmt.Println("num的地址是 ", &num) //0xc00008c000

	var ptr *int = &num
	fmt.Println(ptr) //0xc00008c000

	fmt.Println(&ptr) //0xc0000b4008

	fmt.Println(*ptr) //100

	fmt.Printf("ptr 类型 %T", ptr)
}
