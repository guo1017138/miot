package auth

import (
	"fmt"
	"testing"

	"github.com/luanruisong/miot/consts"
)

func TestLogin(t *testing.T) {
	info := ServiceLoginReq{
		Sid:      consts.SID_XIAOMIIO,
		Username: "user",
		Password: "pass",
	}
	fmt.Println(Login(info), "====")
}
