package qingniao

const (
	// 验证码短信
	smsTypeCode smsType = 1
	// 通知短信
	smsTypeNotify smsType = 2
	// 营销短信
	smsTypeAdvertising smsType = 3
)

// smsType 短信类型
type smsType uint
