package auth

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"

	"github.com/luanruisong/miot/consts"
	"github.com/luanruisong/miot/internal/apis"
	"github.com/luanruisong/miot/internal/token"
	"github.com/luanruisong/miot/internal/utils"
)

func Login(info ServiceLoginReq) error {
	token.ResetToken() // clear old token in case user changed
	resp, err := apis.AuthReq().SetQueryParams(map[string]string{
		"sid":   info.Sid,
		"_json": "true",
	}).Get(apis.AuthURI("/pass/serviceLogin"))
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("resp err:%d", resp.StatusCode())
	}
	ret, err := utils.Decode[ServiceLoginRet](resp.Body())
	if err != nil {
		return err
	}
	tk := token.GetToken()
	if ret.Code != 0 {
		// if passtoken set, use passtoken login
		if info.UserId != 0 && info.DeviceId != "" && info.PassToken != "" {
			fmt.Printf("$$$$$$$$$$$$$$$$$$$: %+v\n", info)
			tk.UserId = info.UserId
			tk.DeviceId = info.DeviceId
			tk.PassToken = info.PassToken
		}
		ret, err = serverLoginAuth(ret, info.Username, info.Password)
		if err != nil {
			return err
		}
		if ret.Code != 0 {
			return fmt.Errorf(ret.Desc)
		}
	}
	tk.UserId = ret.UserId
	tk.PassToken = ret.PassToken
	serviceToken, err := generateServiceToken(ret.Location, ret.Nonce, ret.Ssecurity)
	if err != nil {
		return err
	}
	return tk.SetSubToken(info.Sid, ret.Ssecurity, serviceToken).Sync()
}

func serverLoginAuth(req ServiceLoginRet, user, pass string) (ServiceLoginRet, error) {
	data := map[string]string{
		"_json":    "true",
		"qs":       req.Qs,
		"sid":      req.Sid,
		"_sign":    req.Sign,
		"callback": req.Callback,
		"user":     user,
		"hash":     utils.GetMD5Hash(pass),
	}
	resp, err := apis.AuthReq().SetFormData(data).Post(apis.AuthURI("/pass/serviceLoginAuth2"))
	if err != nil {
		return ServiceLoginRet{}, err
	}
	if !resp.IsSuccess() {
		return ServiceLoginRet{}, fmt.Errorf("resp err:%d", resp.StatusCode())
	}
	return utils.Decode[ServiceLoginRet](resp.Body())
}

func generateServiceToken(location string, nonce int64, ssecurity string) (string, error) {
	nsec := fmt.Sprintf("nonce=%d&%s", nonce, ssecurity)
	hash := sha1.Sum([]byte(nsec))
	encoded := base64.StdEncoding.EncodeToString(hash[:])
	u, _ := url.Parse(location)
	query := u.Query()
	query.Set("clientSign", encoded)
	u.RawQuery = query.Encode()
	resp, err := apis.AuthReq().Get(u.String())
	if err != nil {
		return "", err
	}
	if !resp.IsSuccess() {
		return "", fmt.Errorf("resp err:%d", resp.StatusCode())
	}
	for _, v := range resp.Cookies() {
		if v.Name == "serviceToken" {
			return v.Value, nil
		}
	}
	return "", errors.New("can not find service Token")
}

func AutoLogin(sid string) error {
	if err := token.CheckLogin(sid); err != nil {
		if err = consts.CheckEnv(); err != nil {
			return err
		}
		info := ServiceLoginReq{
			Sid:      sid,
			Username: consts.GetUser(),
			Password: consts.GetPass(),
		}
		return Login(info)
	}
	return nil
}
