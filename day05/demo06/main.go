package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	myMap = make(map[int]int, 10)

	// 互斥锁
	lock sync.Mutex
)

func main() {

	//获取CPU数目
	fmt.Println(runtime.NumCPU())

	//会报错  并发修改map  两种方法
	//1、使用互斥锁解决
	//2、使用goroutine
	for i := 0; i <= 200; i++ {
		go test(i)
	}

}

//goroutine  1-200各个数的阶乘
func test(n int) {
	res := 1
	for i := 1; i < n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}
