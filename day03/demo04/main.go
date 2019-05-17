package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr [26]byte
	for i := 0; i < len(arr); i++ {
		arr[i] = 'A' + byte(i)
		fmt.Printf("%c", arr[i])
	}
	fmt.Println("")
	var numbers [5]int
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(numbers); i++ {
		numbers[i] = rand.Intn(100)
		fmt.Println(numbers[i])
	}

}
