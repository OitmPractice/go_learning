package main

import (
	"fmt"
)

func main() {
	//第一种方式 切片是引用类型，内容改变后，原数组也改变
	var intArr [5]int = [...]int{1, 3, 2, 4, 7}
	slice := intArr[1:3]
	fmt.Println(slice)
	fmt.Println("slice容量=", cap(slice))
	fmt.Println("slice长度=", len(slice))

	//第二种
	var slice2 []int = make([]int, 10)
	slice2[4] = 100
	fmt.Println(slice2)

	var slice3 []string = make([]string, 10)
	slice3[0] = "haha"
	fmt.Println(slice3)

	//第三种
	var slice4 []string = []string{"oitm"}
	slice4 = append(slice4, "momo")
	fmt.Println(slice4)
	slice4 = append(slice4, slice4...)
	fmt.Println(slice4)

	copy(slice4, slice3)
	fmt.Println(slice4)
}
