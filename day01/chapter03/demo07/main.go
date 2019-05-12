package main

import (
	"fmt"
)

// 全局匿名函数
var (
	Func1 = func(a int, b int) int {
		return a + b
	}
)

func main() {
	sum := sum(10, 10)
	fmt.Println(sum)

	// 匿名函数
	result := func(a int, b int) int {
		return a + b
	}(10, 20)
	fmt.Println(result)

	anno := func(a int, b int) int {
		return a + b
	}
	test := anno(30, 30)
	fmt.Println(test)

}

func sum(a int, b int) int {
	return a + b
}

func add(args ...int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
