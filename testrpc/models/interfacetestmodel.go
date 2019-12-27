package models

import "fmt"

type Interfasetestmodel interface {
	Add()
	Subtract()
	Multiply()
	Divide()
}

type interfasetestmodel struct {
	A int
	B int
}

func InitInterfaceModel(a, b int) Interfasetestmodel {
	return &interfasetestmodel{
		a,
		b,
	}
}

func (this *interfasetestmodel) Add() {
	fmt.Println("a + b = ", this.A+this.B)
}

func (this *interfasetestmodel) Subtract() {
	fmt.Println("a - b = ", this.A-this.B)
}

func (this *interfasetestmodel) Multiply() {
	fmt.Println("a * b = ", this.A*this.B)
}

func (this *interfasetestmodel) Divide() {
	fmt.Println("a ÷ b = ", this.A/this.B)
}
