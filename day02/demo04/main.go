package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//生成随机数 我们需要为rand设置一个种子
	rand.Seed(time.Now().Unix())
	n := rand.Intn(100) + 1
	fmt.Println(n)
}
