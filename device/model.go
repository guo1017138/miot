package device

type DeviceExtraSplit struct{}

type DeviceExtra struct {
	IsSetPinCode    int              `json:"isSetPincode"`
	PinCodeType     int              `json:"pincodeType"`
	FwVersion       string           `json:"fw_version"`
	McuVersino      string           `json:"mcu_version"`
	IsSubGroup      bool             `json:"isSubGroup"`
	ShowGroupMember bool             `json:"showGroupMember"`
	Split           DeviceExtraSplit `json:"split"`
}

type DeviceInfo struct {
	Did         string      `json:"did"`
	Token       string      `json:"token"`
	Longitude   string      `json:"longitude"`
	Latitude    string      `json:"latitude"`
	Name        string      `json:"name"`
	Pid         string      `json:"pid"`
	Localip     string      `json:"localip"`
	Mac         string      `json:"mac"`
	Ssid        string      `json:"ssid"`
	Bssid       string      `json:"bssid"`
	ParentId    string      `json:"parent_id"`
	ParentModel string      `json:"parent_model"`
	ShowMode    int         `json:"show_mode"`
	Model       string      `json:"model"`
	AdminFlag   int         `json:"adminFlag"`
	ShareFlag   int         `json:"shareFlag"`
	PermitLevel int         `json:"permitLevel"`
	IsOnline    bool        `json:"isOnline"`
	Desc        string      `json:"desc"`
	Uid         int         `json:"uid"`
	PdId        int         `json:"pd_id"`
	Password    string      `json:"password"`
	P2PId       string      `json:"p2p_id"`
	Rssi        int         `json:"rssi"`
	FamilyId    int         `json:"family_id"`
	ResetFlag   int         `json:"reset_flag"`
	DescNew     string      `json:"desc_new,omitempty"`
	DescTime    []int       `json:"desc_time,omitempty"`
	SpecType    string      `json:"spec_type"`
	Extra       DeviceExtra `json:"extra"`
	OrderTime   int64       `json:"orderTime"`
	FreqFlag    bool        `json:"freqFlag"`
	HideMode    int         `json:"hide_mode"`
	Cnt         int         `json:"cnt"`
	LastOnline  int64       `json:"last_online,omitempty"`
	ComFlag     int         `json:"comFlag"`
	CarPicInfo  string      `json:"carPicInfo"`
}

type DeviceListResult struct {
	List         []DeviceInfo `json:"list"`
	NextStartDid string       `json:"next_start_did"`
	HasMore      bool         `json:"has_more"`
}

type ActionDetail struct {
	Did  string `json:"did"`
	Siid int    `json:"siid"`
	Aiid int    `json:"aiid"`
	In   []any  `json:"in"`
}

type ActionResult struct {
	Did         string `json:"did"`
	Miid        int    `json:"miid"`
	Siid        int    `json:"siid"`
	Aiid        int    `json:"aiid"`
	Code        int    `json:"code"`
	ExeTime     int    `json:"exe_time"`
	NetCost     int    `json:"net_cost"`
	Otlocalts   int64  `json:"otlocalts"`
	OaCost      int    `json:"oa_cost"`
	OaRpcCost   int    `json:"_oa_rpc_cost"`
	WithLatency int    `json:"withLatency"`
}

type GetPropReq struct {
	Did  string `json:"did"`
	Siid int    `json:"siid"`
	Piid int    `json:"piid"`
}

type GetPropResult struct {
	Code       int    `json:"code"`
	Did        string `json:"did"`
	ExeTime    int    `json:"exe_time"`
	Iid        string `json:"iid"`
	Siid       int    `json:"siid"`
	Piid       int    `json:"piid"`
	UpdateTime int    `json:"updateTime"`
	Value      any    `json:"value"`
}

type SetPropReq struct {
	Did   string `json:"did"`
	Siid  int    `json:"siid"`
	Piid  int    `json:"piid"`
	Value any    `json:"value"`
}

type SetPropResult struct {
	Code    int    `json:"code"`
	Did     string `json:"did"`
	ExeTime int    `json:"exe_time"`
	Iid     string `json:"iid"`
	Siid    int    `json:"siid"`
	Piid    int    `json:"piid"`
}
