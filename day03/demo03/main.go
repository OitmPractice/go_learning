package main

import (
	"fmt"
)

func main() {
	//数组：存放同一类型数据，数组在Go中是值类型  未设置值的默认是0
	var arr [6]float64
	arr[0] = 0
	arr[1] = 10
	fmt.Println(arr)

	//数组的地址就是第一个元素的地址
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", &arr[0])

	// var score [5]float64
	// for i := 0; i < len(score); i++ {
	// 	fmt.Printf("请输入第%d个学生成绩", i)
	// 	fmt.Scanln(&score[i])
	// }
	// for i := 0; i < len(score); i++ {
	// 	fmt.Println(score[i])
	// }

	//四种初始化数组方式
	var numArr02 = [3]int{1, 2, 3}
	fmt.Println(numArr02)

	var numArr03 [3]int = [3]int{1, 2, 3}
	fmt.Println(numArr03)

	var numArr04 = [...]int{1, 2, 3}
	fmt.Println(numArr04)

	numArr05 := [...]int{1: 100, 0: 30, 5: 9}
	fmt.Println(numArr05)
}
