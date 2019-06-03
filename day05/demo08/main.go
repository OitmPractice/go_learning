package main

import (
	"fmt"
)

func main() {
	var intChan = make(chan int, 50)
	var exitChan = make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}

func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Printf("writeData 写入了数据 %v \n", i)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("redaData 读到的数据=%v \n", v)
	}

	exitChan <- true
	close(exitChan)
}
