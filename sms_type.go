package qingniao

const (
	// 普通短信
	smsTypeCommon smsType = iota
	// 通知短信
	smsTypeNotify
	// 营销短信
	smsTypeAdvertising
)

// smsType 短信类型
type smsType uint
