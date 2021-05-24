package una

type sms struct {
	// 短信类型
	smsType SmsType `validate:"required,oneof=1 2 3"`
	// 手机号列表
	mobiles []string `validate:"required,dive,mobile"`
}
