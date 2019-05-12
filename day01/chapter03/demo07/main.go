package main

import (
	"fmt"
)

func main() {
	sum := sum(10, 10)
	fmt.Println(sum)

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
