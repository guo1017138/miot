package apis

import (
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/go-resty/resty/v2"
	"github.com/luanruisong/miot/consts"
	"github.com/luanruisong/miot/internal/token"
)

func AuthReq() *resty.Request {
	app := token.GetToken()
	cookies := append([]*http.Cookie{}, &http.Cookie{
		Name:  "sdkVersion",
		Value: "3.9",
	}, &http.Cookie{
		Name:  "deviceId",
		Value: app.DeviceId,
	})
	if app.IsLogin() {
		cookies = append(cookies, &http.Cookie{
			Name:  "userId",
			Value: fmt.Sprintf("%d", app.UserId),
		}, &http.Cookie{
			Name:  "passToken",
			Value: app.PassToken,
		})
	}
	header := map[string]string{
		"User-Agent": "APP/com.xiaomi.mihome APPV/6.0.103 iosPassportSDK/3.9.0 iOS/14.4 miHSTS",
	}
	req := resty.New()
	return req.R().SetCookies(cookies).SetHeaders(header)
}

func AppReq(sid string) *resty.Request {
	if err := token.CheckLogin(sid); err != nil {
		panic(err)
	}
	tk := token.GetToken()
	cookies := append([]*http.Cookie{}, &http.Cookie{
		Name:  "PassportDeviceId",
		Value: tk.DeviceId,
	}, &http.Cookie{
		Name:  "userId",
		Value: fmt.Sprintf("%d", tk.UserId),
	}, &http.Cookie{
		Name:  "serviceToken",
		Value: tk.GetSubToken(sid).ServiceToken,
	})
	header := map[string]string{
		"User-Agent":                 "APP/com.xiaomi.mihome APPV/6.0.103 iosPassportSDK/3.9.0 iOS/14.4 miHSTS",
		"x-xiaomi-protocal-flag-cli": "PROTOCAL-HTTP2",
	}
	req := resty.New()
	return req.R().SetCookies(cookies).SetHeaders(header)
}

func AuthURI(uri string) string {
	return _uri(consts.AuthHost, uri)
}

func AppURI(uri string) string {
	return _uri(consts.AppHost, uri)
}

func _uri(host, uri string) string {
	u, _ := url.Parse(host)
	u.Path = path.Join(u.Path, uri)
	return u.String()
}
