package interfacetestmodel

import "fmt"

type Binterface interface {
	Witer()
}

type binterface struct {
	CommonElements string //公共元素
}

func InitBF(commonElements string) Binterface {
	return binterface{CommonElements: commonElements}
}

func (this binterface) Witer() {
	fmt.Println("我在写：", this.CommonElements)
}
