package auth

type ServiceLoginReq struct {
	Sid       string `json:"sid"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	UserId    int    `json:"user_id"`
	DeviceId  string `json:"device_id"`
	PassToken string `json:"passToken"`
}

type ServiceLoginRet struct {
	Code      int    `json:"code"`
	Desc      string `json:"desc"`
	Qs        string `json:"qs"`
	Sid       string `json:"sid"`
	Sign      string `json:"_sign"`
	Callback  string `json:"callback"`
	PassToken string `json:"passToken"`
	UserId    int    `json:"userId"`
	Location  string `json:"location"`
	Nonce     int64  `json:"nonce"`
	Ssecurity string `json:"ssecurity"`
}
