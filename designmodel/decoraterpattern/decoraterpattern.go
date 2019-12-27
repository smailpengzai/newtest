package main

import (
	"fmt"
)

type Company interface {
	Showing()
}

type BaseCompany struct {
}

func (pB *BaseCompany) Showing() {
	fmt.Println("公司有老板，有前台，有人事...")
}

type DevelopingCompany struct {
	Company
}

func (pD *DevelopingCompany) AddWorker() {
	fmt.Println("还有开发、测试、财务人员")
}

func (pD *DevelopingCompany) Showing() {
	fmt.Println("发展中公司中：")
	pD.Company.Showing()
	pD.AddWorker()
}

type BigCompany struct {
	Company
}

func (pD *BigCompany) AddWorker() {
	fmt.Println("除此之外，个职能人员应有尽有")
}

func (pD *BigCompany) Showing() {
	fmt.Println("大型公司中：")
	pD.Company.Showing()
	pD.AddWorker()
}

func main() {
	company := &BaseCompany{}
	company.Showing()
	developingCompany := &DevelopingCompany{Company: company}
	developingCompany.Showing()
	bigCompany := &BigCompany{Company: developingCompany}
	bigCompany.Showing()
	return
}
