package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("相机开始工作")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

type Phone struct {
}

func (c Phone) Start() {
	fmt.Println("手机开始工作")
}
func (c Phone) Stop() {
	fmt.Println("手机停止工作")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {

	var usbs [3]Usb
	usbs[0] = Camera{}
	usbs[1] = Phone{}
	fmt.Println(usbs)
	usbs[0].Start()
	usbs[1].Start()

	c := Computer{}
	c.Working(Phone{})
	c.Working(Camera{})
}
