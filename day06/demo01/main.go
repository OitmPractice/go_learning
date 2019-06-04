package main

import (
	"fmt"
	"reflect"
)

func main() {

	const (
		a = iota
		b
		c, d = iota, iota
	)

	fmt.Println(a, b, c, d) //0 1 2 2

	var num int = 100
	reflectTest01(num)

}

func reflectTest01(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType = ", rType)

	rVal := reflect.ValueOf(b)
	fmt.Println("rVal = ", rVal)
	fmt.Printf("rVal =  %T \n", rVal) // rVal =  reflect.Value

	n2 := 2 + rVal.Int()
	fmt.Println(n2)

	iV := rVal.Interface()

	num2 := iV.(int)
	fmt.Println(num2)
}
