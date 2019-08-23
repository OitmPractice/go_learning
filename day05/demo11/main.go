package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var (
	pool *redis.Pool
)

func main() {
	testRedisPool()
}

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   // 最大空闲连接数
		MaxActive:   0,   //表示和数据库的最大连接数 ， 0表示没有限制
		IdleTimeout: 100, // 最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化连接的代码
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

// redis连接池
func testRedisPool() {

	conn := pool.Get()
	defer conn.Close()

	v, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
}

func testHash() {
	//连接redis  获得conn
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("conn == ", conn)

	defer conn.Close()

	//向redis写入数据
	_, err = conn.Do("HSet", "user", "name", "Oitm")
	// _, err = conn.Do("HMSet", "user", "name", "Oitm", "age", 18)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("操作ok")

	//获取
	//利用redis提供的方法进行string转换
	r, err := redis.String(conn.Do("HGet", "user", "name"))
	// r, err := redis.Strings(conn.Do("HMGet", "user", "name", "age"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)

}

func testRedisString() {
	//连接redis  获得conn
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("conn == ", conn)

	defer conn.Close()

	//向redis写入数据
	_, err = conn.Do("set", "name", "Oitm")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("操作ok")

	//获取
	//利用redis提供的方法进行string转换
	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}
