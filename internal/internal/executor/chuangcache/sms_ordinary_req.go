package chuangcache

// SmsOrdinaryReq 发送普通短信请求
type SmsOrdinaryReq struct {
	*SmsBaseReq

	// 短信类型
	// 默认：1
	// 通知短信：2
	Type int `json:"sms_type"`
}
