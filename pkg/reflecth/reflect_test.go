package reflecth

import (
	"fmt"
	"testing"
)

type demoTest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
	ID       int64  `json:"_id"`
	NickName string `json:"nick_name"`
}

func TestGetStructValeByField(t *testing.T) {
	v := demoTest{
		Username: "cc",
	}

	name := GetStructValeByField(&v, "Username")
	str := fmt.Sprintf("%v", name)
	if str != "cc" {
		t.Errorf("GetStructValeByField faild %#v  %#v", name, str)
	}
}

func TestSetStructValeByField(t *testing.T) {
	v := demoTest{
		Username: "a",
	}

	err := SetStructValeByField(&v, "Username", "bb")
	if err != nil {
		t.Errorf("SetStructValeByField get err %#v", err)
		return
	}

	if v.Username != "bb" {
		t.Errorf("SetStructValeByField faild %#v", v)
		return
	}
}

func TestGetName(t *testing.T) {
	v := &demoTest{}
	if GetName(v) != "demoTest" {
		t.Errorf("GetName faild %#v", v)
		return
	}
}
