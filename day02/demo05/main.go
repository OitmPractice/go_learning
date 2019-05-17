package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// len函数长度按照字节计算的
	str := "OOOOitm默默"
	fmt.Println("len = ", len(str)) // = 10

	//处理中文需要把字符串 转化为rune的切片
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c \n", str2[i])
	}

	//str和int转换
	n, err := strconv.Atoi("12")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}
	str3 := strconv.Itoa(12345)
	fmt.Printf("str3 = %v, type = %T \n", str3, str3)

	//字符串和字节互转
	bytes := []byte(str)
	fmt.Printf("bytes = %v, bytes = %T \n", bytes, bytes)

	str4 := string(bytes)
	fmt.Println(str4)

	//进制之间转换 二进制
	str5 := strconv.FormatInt(16, 2)
	fmt.Println(str5)

	//字符串包含
	b := strings.Contains(str, "O")
	fmt.Println(b)

	count := strings.Count(str, "O")
	fmt.Println(count)

	//不区分大小写
	b2 := strings.EqualFold("A", "a")
	fmt.Println(b2)
	fmt.Println("abc" == "ABC") //字符串比较

	b3 := strings.HasPrefix(str, "O")
	fmt.Println("b3=", b3)

	index := strings.Index(str, "tm") // 字符串包含字符的位置
	fmt.Println(index)

	index2 := strings.LastIndex(str, "O") // 字符串最后出现字符的位置
	fmt.Println(index2)

	//字符串替换 最后一个参数1只替换一个，-1替换所有
	str6 := strings.Replace(str, "O", "N", -1)
	fmt.Println(str6)

	//字符串分割
	strArr := strings.Split(str, "i")
	fmt.Println(strArr)

	str7 := strings.ToLower(str)
	fmt.Println(str7)

	//去掉左右空格
	str8 := strings.TrimSpace("  oitm hen shuai   ")
	fmt.Println(str8)
	//去掉左右空格+指定字符
	str9 := strings.Trim(str8, "o")
	fmt.Println(str9)

	//string切片处理
	slice := str9[4:]
	fmt.Println(slice)

}
