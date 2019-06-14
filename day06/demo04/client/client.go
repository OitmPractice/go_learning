package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("err == ", err)
	}
	fmt.Println("conn success = ", conn)

	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取失败 err == ", err)
	}
	// n是字节
	n, err := conn.Write([]byte(str))
	if err != nil {
		fmt.Println("发送失败 err == ", err)
	}
	fmt.Println(n)

}
