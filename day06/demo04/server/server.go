package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//循环接收客户端数据
	defer conn.Close()

	for {
		fmt.Println("server waiting for wirte!!!")
		buf := make([]byte, 1024)
		//一直等待客户端发送信息，客户端没有write 会一直等待。
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err == ", err)
			return
		} else {
			fmt.Println(string(buf[:n]))
		}
	}

}

func main() {
	fmt.Println("服务器开始监听")

	// tcp：表示网络协议
	listen, err := net.Listen("tcp", "0.0.0.0:8888")

	if err != nil {
		fmt.Println("err == ", err)
		return
	}

	defer listen.Close()

	//保证一直监听
	for {
		// 一直等待客户端连接
		fmt.Println("等待连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept() 出错了 ", err)
		} else {
			fmt.Printf("accept() 成功了 conn = %v \n ", conn)
			fmt.Println(conn.RemoteAddr())
			//准备起一个协程 为客户端服务
			go process(conn)
		}
	}
	fmt.Println("listen success == ", listen)
}
