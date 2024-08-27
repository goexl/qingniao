package chuangcache

// SmsRsp 发送短信响应
type SmsRsp struct {
	// 接口返回码
	Code int `json:"code"`
	// 发送短信流水号
	Id string `json:"sendid"`
	// 接口返回信息
	Msg string `json:"msg"`
}
