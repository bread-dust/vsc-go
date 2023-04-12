package simplefactory

import "fmt"

// API is interface
type API interface{
	Say(name string) string
}

// NewAPI return API instance by type
func NewAPI(t int) API{
	if t==1{
		return &hiAPI{}
	} else if t ==2 {
		return &helloAPI{}
	}
	return nil
}

// hiPAI is a implement of API
type hiAPI struct{}

// Say hi to name
func (h *hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi,%s",name)
}

// helloAPI is a implement of API
type helloAPI struct{
}

// S hello t oname
func (he *helloAPI) Say(name string) string {
	return fmt.Sprintf("hello to %s",name)
}

