package main

import (
	"fmt"
	"go_code/day05/demo01/service"
)

type customerView struct {
	//定义必要属性
	key             string //接收用户输入
	loop            bool
	customerService *service.CustomerService
}

func (this *customerView) list() {
	customers := this.customerService.List()
	fmt.Println("--------------客户列表--------------")

	for i := 0; i < len(customers); i++ {

	}

	fmt.Println("--------------客户列表完成--------------")
}

func (this *customerView) mainMenu() {
	for {
		fmt.Println("客户信息管理软件")
		fmt.Println("1、添加客户")
		fmt.Println("2、修改客户")
		fmt.Println("3、删除客户")
		fmt.Println("4、客户列表")
		fmt.Println("5、退出")
		fmt.Println("请选择1-5")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			fmt.Println("添加客户")
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			fmt.Println("客户列表")
		case "5":
			this.loop = false
		default:
			fmt.Println("您的输入有误，请重新输入")
		}
		if this.loop == false {
			break
		}
	}
	fmt.Println("您退出了客户关系管理系统！")
}

func main() {

	//如何删除切片元素
	var slice []int = []int{1, 2, 3, 4, 5}
	slice = append(slice[:2], slice[3:]...)
	fmt.Println(slice)

	customerView := customerView{
		key:  "",
		loop: true,
	}
	customerView.customerService = service.NewCustomerService()

	customerView.mainMenu()
}
