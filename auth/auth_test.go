package auth

import (
	"fmt"
	"github.com/luanruisong/miot/utils"
	"testing"
)

func TestLogin(t *testing.T) {
	fmt.Println(Login(utils.SID_XIAOMIIO, "user", "pass"), "====")
}