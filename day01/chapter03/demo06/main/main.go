package main

import (
	"fmt"
	"go_code/day01/chapter03/demo06/model"
)

func main() {
	fmt.Println(model.HeroName)

label1:
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 1 {
				break label1
			}
		}
	}
}
