package main

import (
	"fmt"
)

func main() {
	//类型推断
	var a interface{}
	var b float32 = 1.23
	a = b
	x, ok := a.(float32)
	if ok {
		fmt.Println(x)
	}

	number := 100
	TypeJudge(number)

}

func TypeJudge(items ...interface{}) {
	for _, item := range items {
		switch item.(type) {
		case bool:
			fmt.Printf("item is %T\n", item)
		case int:
			fmt.Printf("item is %T\n", item)
		default:
			fmt.Printf("item is %T\n", item)
		}
	}
}
