package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		// fmt.Println(i)
	}

	k := 10
	for {
		fmt.Println("Oitm")
		if k <= 5 {
			break
		}
		k--
	}

	var str = "Oitm"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}

	// 默认使用字符方式遍历  对中文支持很好 推荐使用
	for index, val := range str {
		fmt.Printf("%d %c\n", index, val)
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 0; i < 10; i++ {
		for k := 0; k < (10-i)/2; k++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

} 
