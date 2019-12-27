package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

type Company interface {
	Runing()
}

type BigCompany struct {
	Worker
}

func (pB *BigCompany) Running() {
	fmt.Println("员工都是螺丝钉")
	pB.Worker.Leaving()
	fmt.Println("员工跑路后随时可以找人顶替")
}

type SmallCompany struct {
	Worker
}

func (pS *SmallCompany) Running() {
	fmt.Println("随便一个员工都是骨干")
	pS.Worker.Leaving()
	fmt.Println("员工跑路后，公司运转受阻")
}

type Worker interface {
	Leaving()
}

type GoodWorker struct {
}

func (pG *GoodWorker) Leaving() {
	fmt.Println("好员工跑路后")
}

type NormalWorker struct {
}

func (pN *NormalWorker) Leaving() {
	fmt.Println("普通员工跑路后")
}

func main() {
	//员工的本质是不变的
	pgoodworker := &GoodWorker{}
	pnormalworker := &NormalWorker{}
	//在不同的公司不同的员工的作用
	beego.Critical("大公司 的好员工跑路后")
	pbigCompany := &BigCompany{Worker: pgoodworker}
	pbigCompany.Running()
	beego.Critical("大公司 的普通员工跑路后")
	pbigCompany2 := &BigCompany{Worker: pnormalworker}
	pbigCompany2.Running()
	beego.Critical("小公司 的好员工跑路后")
	psmallCompany := &SmallCompany{Worker: pgoodworker}
	psmallCompany.Running()
	beego.Critical("小公司 的普通员工跑路后")
	psmallCompany2 := &SmallCompany{Worker: pnormalworker}
	psmallCompany2.Running()

	return
}
