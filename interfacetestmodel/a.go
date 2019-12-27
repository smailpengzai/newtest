package interfacetestmodel

import "fmt"

type Ainterface interface {
	Read()
}

type ainterface struct {
	CommonElements string //公共元素
}

func InitAF(commonElements string) Ainterface {
	return ainterface{CommonElements: commonElements}
}

func (this ainterface) Read() {
	fmt.Println("我在读：", this.CommonElements)
}
