package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	//返回的file是一个指针
	file, err := os.Open("./temp.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("file == %v \n", file)

	defer file.Close()

	//创建一个 *Reader  是带缓冲的  一次一次的读
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}

	//一次性读取文件  小文件使用
	content, err := ioutil.ReadFile("./temp.txt")
	if err == nil {
		fmt.Println(string(content))
	} else {
		fmt.Println(err)
	}

}
