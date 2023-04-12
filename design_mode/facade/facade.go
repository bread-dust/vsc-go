package facade

import "fmt"

// 股民 - 基金 - 股票

// API is a facade interface of facade package
type API interface{
	Test() string
}


// apiImpl is a facade implement
type apiImpl struct{
	a AModuleAPI
	b BModuleAPI
}

// NewAPI in a instancer of API interface
func NewAPI()API{
	return &apiImpl{
		a:NewAModuleAPI(),
		b:NewBModuleAPI(),
	}
}

func (a *apiImpl) Test() string{
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s",aRet,bRet)
}

// NewAModuleAPI return new AModuleAPI
func NewAModuleAPI()AModuleAPI{
	return &aModuleImpl{}
}

type AModuleAPI interface{
	TestA() string
}

type aModuleImpl struct{}

func (a*aModuleImpl)TestA()string{
	return "A module is running"
}

// NewBModuleAPI return new AModuleAPI
func NewBModuleAPI()BModuleAPI{
	return &bModuleImpl{}
}

type BModuleAPI interface{
	TestB() string
}

type bModuleImpl struct{}

func (b*bModuleImpl)TestB()string{
	return "B module is running"
}