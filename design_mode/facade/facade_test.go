package facade

import "testing"

var expect = "A module is running\nB module is running"

// TestFacadeAPI 测试
func TestFacadeAPI(t *testing.T){
	api:= NewAPI()
	ret := api.Test()
	if ret!=expect{
		t.Fatalf("expect %s return %s",expect,ret)
	}
}