package main

import (
	"fmt"
)

func main() {
	var isSelect = false
	if isSelect {
		fmt.Println("momo")
	} else if age := 20; age > 18 {
		fmt.Println("Oitm")
	}

	var age = 10
	switch age {
	case 1:
		fmt.Println(age)
	default:
		fmt.Println(age)
	}

}
