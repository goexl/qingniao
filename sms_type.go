package qingniao

const (
	// SmsTypeCommon 普通短信
	SmsTypeCommon SmsType = 1
	// SmsTypeNotify 通知短信
	SmsTypeNotify SmsType = 2
	// SmsTypeAdvertising 营销短信
	SmsTypeAdvertising SmsType = 3
)

// SmsType 短信类型
type SmsType int
