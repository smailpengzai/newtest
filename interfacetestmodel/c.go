package interfacetestmodel

type Cinterface struct {
	CommonElements string //公共元素

	Ainterface Ainterface
	Binterface Binterface
}

func Total() {
	a := "fdsafdsaf"
	//var cinterface Cinterface
	//cinterface.CommonElements = "ljsljdflksf"
	//cinterface.Binterface = InitBF(cinterface.CommonElements)
	//cinterface.Ainterface = InitAF(cinterface.CommonElements)

	//cinterface.Ainterface.Read()
	//cinterface.Binterface.Witer()

	cinterface := initAll(a)

	cinterface.Binterface.Witer()
	cinterface.Ainterface.Read()
}

func initAll(commonElements string) (cinterface Cinterface) {
	cinterface.CommonElements = commonElements
	cinterface.Binterface = InitBF(cinterface.CommonElements)
	cinterface.Ainterface = InitAF(cinterface.CommonElements)
	return
}
