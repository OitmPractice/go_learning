package main

import (
	"fmt"
)

func main() {
	channelCount := 8
	intChan := make(chan int, 10000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, channelCount)

	go putNum(intChan)
	for i := 0; i < 8; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < channelCount; i++ {
			//这里会阻塞
			<-exitChan
		}
		close(primeChan)
		close(exitChan)
	}()

	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数=%d\n", res)
	}
}



//存入1-8000的数
func putNum(intChan chan int) {
	for i := 1; i <= 800; i++ {
		intChan <- i
	}
	close(intChan)
}

//判断是否是素数
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	flag := false

	for {
		num, ok := <-intChan
		if !ok {
			break
		}

		flag = false
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = true
				break
			}
		}

		if !flag {
			primeChan <- num
		}
	}

	fmt.Println("有一个primeNum 协程因为取不到数据，退出")
	exitChan <- true
}
