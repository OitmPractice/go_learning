package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//协程一般以函数为单位开始的
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("Main hello golang ", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello world! " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
