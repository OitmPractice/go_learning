package main

import (
	"fmt"
)

func main() {

	num := new(int)
	fmt.Println(num)  //0xc000016090
	fmt.Println(*num) //0

}
