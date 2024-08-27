package constant

const (
	// 验证码短信
	SmsTypeCode SmsType = 1
	// 通知短信
	SmsTypeNotify SmsType = 2
	// 营销短信
	SmsTypeAdvertising SmsType = 3
)

// SmsType 短信类型
type SmsType uint
