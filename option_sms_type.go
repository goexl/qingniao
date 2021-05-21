package una

var _ option = (*optionSmsType)(nil)

type optionSmsType struct {
	smsType SmsType
}

// CommonSms 配置普通短信
func CommonSms() *optionSmsType {
	return &optionSmsType{
		smsType: SmsTypeCommon,
	}
}

// CommonNotify 配置通知短信
func CommonNotify() *optionSmsType {
	return &optionSmsType{
		smsType: SmsTypeNotify,
	}
}

// CommonAdvertising 配置营销（广告）短信
func CommonAdvertising() *optionSmsType {
	return &optionSmsType{
		smsType: SmsTypeAdvertising,
	}
}

func (b *optionSmsType) apply(options *options) {
	options.smsType = b.smsType
}
