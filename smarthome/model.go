package smarthome

type GetDeviceDataReq struct {
	Did       string `json:"did"`
	Key       string `json:"key,omitempty"`   // 属性或事件名，可选参数。(注意：如果设备是蓝牙设备，传入的是object id， 且为十进制数据；如果是wifi设备，才传入自定义属性或事件名，可以在开发者平台-产品-功能定义中查看)，如果是miot-spec设备，请传入（siid.piid或者siid.eiid）
	Type      string `json:"type"`            // 必选参数[prop/event], 如果是查询上报的属性则type为prop，查询上报的事件则type为event,
	TimeStart int64  `json:"time_start"`      // 数据起点，单位是秒。必选参数
	TimeEnd   int64  `json:"time_end"`        // 数据终点，单位是秒。必选参数，time_end必须大于time_start,
	Group     string `json:"group,omitempty"` // 返回数据的方式，默认raw,可选值为hour、day、week、month。可选参数.
	Limit     int    `json:"limit,omitempty"` // 返回数据的条数，默认20，最大1000。可选参数.
	Uid       int64  `json:"uid,omitempty"`   // 要查询的用户id 。可选参数
}

type GetDeviceDataResult struct {
}

type GetUserStatisticsReq struct {
	Did       string `json:"did"`
	Key       string `json:"key,omitempty"` // 属性或事件名，可选参数。(注意：如果设备是蓝牙设备，传入的是object id， 且为十进制数据；如果是wifi设备，才传入自定义属性或事件名，可以在开发者平台-产品-功能定义中查看)，如果是miot-spec设备，请传入（siid.piid或者siid.eiid）
	DataType  string `json:"data_type"`     // 数据类型 包括： 采样统计 日统计:stat_day_v3 / 周统计:stat_week_v3 / 月统计:stat_month_v3;
	TimeStart int64  `json:"time_start"`    // 数据起点，单位是秒。必选参数
	TimeEnd   int64  `json:"time_end"`      // 数据终点，单位是秒。必选参数，time_end必须大于time_start,
	Limit     int    `json:"limit"`         // 限制次数，0为默认条数
}

type GetUserStatisticsResult struct {
	Value string `json:"value"`
	Time  int64  `json:"time"`
}
