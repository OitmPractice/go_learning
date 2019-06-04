package main

import (
	"fmt"
	"reflect"
)

func reflect01(v interface{}) {
	rval := reflect.ValueOf(v)

	fmt.Printf("rval的类型是 = %T \n", rval)

	//Elem返回持有的接口保管的值的Value的封装，或者v持有的指针指向的值的Value的封装。如果v的Kind不是interface或Ptr会panic
	rval.Elem().SetInt(20)
}

func main() {
	var num int = 100
	reflect01(&num)
	fmt.Println(num)
}
