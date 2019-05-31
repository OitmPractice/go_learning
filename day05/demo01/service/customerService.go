package service

import (
	"go_code/day05/demo01/model"
)

//完成customer的操作  包括增删改查
type CustomerService struct {
	customers []model.Customer
	//声明当前切片含有多少个客户
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "Oitm", "男", 20, "112", "example@email.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

func (this *CustomerService) List() []model.Customer {
	return this.customers
}
