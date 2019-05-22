package main

import (
	"fmt"
)

func main() {
	//第一种
	//map的声明不会分配内存，需要make为map分配内存空间
	var map1 map[string]string
	map1 = make(map[string]string, 10)
	map1["name"] = "Oitm"
	fmt.Println(map1)

	//第二种
	var map2 = make(map[string]int, 10)
	map2["age"] = 18
	fmt.Println(map2)

	//第三种
	var map3 map[string]string = map[string]string{
		"name": "oitm",
		"age":  "18",
	}
	delete(map3, "age") //删除元素
	fmt.Println(map3)

	//查看某个key是否存在
	val, ok := map3["name"]
	if ok {
		fmt.Println(val)
	}

	//map遍历
	for k, v := range map3 {
		fmt.Println(k, v)

	}
}
