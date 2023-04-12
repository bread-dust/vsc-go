package simplefactory

import "testing"

func Test_hiAPI_Say(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Tom")
	if s != "Hi,Tom" {
		t.Fatal("Type1 test fail")
	}
}

func Test_helloAPI(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Tom")
	if s != "hello to Tom" {
		t.Fatal("Type2 test fail")
	}

}
