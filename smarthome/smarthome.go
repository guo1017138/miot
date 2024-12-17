package smarthome

import (
	"encoding/json"

	"github.com/luanruisong/miot/consts"
	"github.com/luanruisong/miot/internal/apis"
)

/**
 * 查询用户名下设备上报的属性和事件
 * 获取设备属性和事件历史记录，订阅消息直接写入到服务器，不需要插件添加，最多查询90天前的记录。
 * 通下面的set_user_device_data的参数一一对应， /user/get_user_device_data
 * 对于蓝牙设备，params.key 可参考文档 [米家BLE Object定义](https://iot.mi.com/new/doc/embedded-development/ble/object-definition.html)
 *
 * error code:
 *
 * | code | desc |
 * | :-: | --- |
 * |  0  | 成功 |
 * | -8  | 请求参数缺失或者类型不对 |
 * | -4  | 服务器错误 |
 * | -1  | 请求uid无权限获取did的相关数据 |
 *
 * @param {json} params -参数\{did,type,key,time_start,time_end,limit}含义如下：
 * @param {string} params.did 设备id。 必选参数
 * @param {string} params.key 属性或事件名，可选参数。(注意：如果设备是蓝牙设备，传入的是object id， 且为十进制数据；如果是wifi设备，才传入自定义属性或事件名，可以在开发者平台-产品-功能定义中查看)，如果是miot-spec设备，请传入（siid.piid或者siid.eiid）
 * @param {string} params.type 必选参数[prop/event], 如果是查询上报的属性则type为prop，查询上报的事件则type为event,
 * @param {number} params.time_start 数据起点，单位是秒。必选参数
 * @param {number} params.time_end 数据终点，单位是秒。必选参数，time_end必须大于time_start,
 * @param {string} params.group 返回数据的方式，默认raw,可选值为hour、day、week、month。可选参数.
 * @param {string} params.limit 返回数据的条数，默认20，最大1000。可选参数.
 * @param {number} params.uid 要查询的用户id 。可选参数
 * @returns {Promise}
 */
func GetDeviceData(req GetDeviceDataReq) (resp []GetDeviceDataResult, err error) {
	r, err := apis.SignAppPost[[]GetDeviceDataResult](consts.SID_XIAOMIIO, "/user/get_user_device_data", req)
	if err != nil {
		return resp, err
	}
	for _, v := range r {
		json.Unmarshal([]byte(v.Value), &(v.Values))
		resp = append(resp, v)
	}
	return resp, err
}

/*
*
  - 提供返回设备数据统计服务，使用该接口需要配置产品model以支持使用，建议找对接的产品人员进行操作。
  - 图表📈统计接口 /v2/user/statistics
  - 注:由于sds限额问题，可能会出现一次拉不到或者拉不完数据的情况，会返回code:0和message:“sds throttle”
  - @param {object} params
  - @param {string} params.did did
  - @param {string} params.data_type 数据类型 包括： 采样统计 日统计:stat_day_v3 / 周统计:stat_week_v3 / 月统计:stat_month_v3;
  - @param {string} params.key 需要统计的字段，即统计上报对应的key
  - @param {number} params.time_start 开始时间
  - @param {number} params.time_end 结束时间
  - @param {number} params.limit 限制次数，0为默认条数
  - @return {Promise<Object>}
    {
    "code": 0,
    "message": "ok",
    "result": [
    {
    "value": "[12,34]", // 为一个数组形式json串
    "time": 1543593600 // 时间戳
    },
    {
    "value": "[10,11]",
    "time": 1541001600
    }]
    }
*/
func GetUserStatistics(req GetUserStatisticsReq) ([]GetUserStatisticsResult, error) {
	return apis.SignAppPost[[]GetUserStatisticsResult](consts.SID_XIAOMIIO, "/v2/user/statistics", req)
}
