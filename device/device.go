package device

import (
	"github.com/luanruisong/miot/consts"
	"github.com/luanruisong/miot/internal/apis"
)

func List(getVirtualModel bool, getHuamiDevices int) (DeviceListResult, error) {
	return apis.SignAppPost[DeviceListResult](consts.SID_XIAOMIIO, "/home/device_list", map[string]any{
		"getVirtualModel":    getVirtualModel,
		"getHuamiDevices":    getHuamiDevices,
		"get_split_device":   true,
		"support_smart_home": true,
		"get_cariot_device":  true,
		"get_third_device":   true,
	})
}

func Action(action ActionDetail) (ActionResult, error) {
	return apis.SignAppPost[ActionResult](consts.SID_XIAOMIIO, "/miotspec/action", map[string]any{
		"params": action,
	})
}

func GetProp(req GetPropReq) (GetPropResult, error) {
	results, err := GetProps([]GetPropReq{req})
	if err != nil {
		return results[0], err
	}
	return GetPropResult{}, err
}

func GetProps(req []GetPropReq) ([]GetPropResult, error) {
	return apis.SignAppPost[[]GetPropResult](consts.SID_XIAOMIIO, "/miotspec/prop/get", map[string]any{
		"params": req,
	})
}

func SetProp(req SetPropReq) (SetPropResult, error) {
	results, err := SetProps([]SetPropReq{req})
	if err != nil {
		return results[0], err
	}
	return SetPropResult{}, err
}

func SetProps(req []SetPropReq) ([]SetPropResult, error) {
	return apis.SignAppPost[[]SetPropResult](consts.SID_XIAOMIIO, "/miotspec/prop/set", map[string]any{
		"params": req,
	})
}
