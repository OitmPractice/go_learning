package main

import (
	"fmt"
	"time"
)

func main() {

	testRecover()
	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}
}

func testRecover() {
	go sayHello()
	go test()
}

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("发生错误了 err = ", err)
		}
	}()
	var myMap map[int]string
	myMap[0] = "golang"
}

// ----------------------------------------------------------------------------------------------------

func testSelect() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello " + fmt.Sprintf("%d \n", i)
	}

	//传统方法在遍历管道时，如果不关闭会阻塞而导致deadlock
	//在实际开发中，我们可能不去定什么时候关闭管道
	//可以使用select方式解决

	for {
		select {
		case v := <-intChan:
			fmt.Printf("读取到的数据%d\n", v)
			time.Sleep(time.Second)
		case v := <-stringChan:
			fmt.Printf("读取到stringChan = %s \n", v)
		default:
			fmt.Println("都取不到，程序员自己处理")
			break
		}
	}
}
