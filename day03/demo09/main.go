package main

import (
	"fmt"
	"sort"
)

func main() {
	monster := make(map[string]string, 1)
	monster["name"] = "oitm"
	monster["age"] = "18"

	//切片append
	monsters := make([]map[string]string, 0)
	monsters = append(monsters, monster)
	fmt.Println(monsters)

	//map根据key排序
	numbers := make(map[int]int, 10)
	numbers[1] = 20
	numbers[2] = 9

	var keys []int
	for k, _ := range numbers {
		keys = append(keys, k)
	}
	sort.Ints((keys))
	for _, k := range keys {
		fmt.Println(numbers[k])
	}

}
