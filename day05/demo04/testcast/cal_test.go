package main

import (
	"fmt"
	"testing"
)

func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 45 {
		t.Fatalf("程序执行错误，期望值=%v  实际值=%v", 55, res)
	}
	t.Logf("addUpper(10)执行正确")
}

func TestHello(t *testing.T) {
	fmt.Println("调用了TestHello")
}
