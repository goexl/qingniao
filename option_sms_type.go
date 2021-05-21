package una

var _ option = (*optionSmsType)(nil)

type optionSmsType struct {
	sms sms
}

// CommonSms 配置普通短信
func CommonSms(target string) *optionSmsType {
	return &optionSmsType{
		sms: sms{
			Type:    SmsTypeCommon,
			targets: []string{target},
		},
	}
}

// NotifySms 配置通知短信
func NotifySms(target string) *optionSmsType {
	return &optionSmsType{
		sms: sms{
			Type:    SmsTypeNotify,
			targets: []string{target},
		},
	}
}

// AdvertisingSms 配置营销（广告）短信
func AdvertisingSms(targets ...string) *optionSmsType {
	return &optionSmsType{
		sms: sms{
			Type:    SmsTypeAdvertising,
			targets: targets,
		},
	}
}

func (b *optionSmsType) apply(options *options) {
	options.sms = b.sms
}
