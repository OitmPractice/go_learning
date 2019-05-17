package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 10)
	slice = []int{20, 10, 30, 2, 60}
	fmt.Println(slice)

	bubboSort(&slice)
	fmt.Println("slice2 = ", slice)
}

func bubboSort(arr *[]int) {
	fmt.Println(*arr)

	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i; j++ {
			if (*arr)[j] > (*arr)[j+i] {
				temp := (*arr)[j+1]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
			fmt.Println(j)
		}
	}
	fmt.Println("*arr = ", *arr)
}
