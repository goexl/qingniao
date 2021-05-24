package una

var _ option = (*optionSms)(nil)

type optionSms struct {
	smsType SmsType
	mobiles []string
}

// CommonSms 配置普通短信
func CommonSms(target string) *optionSms {
	return &optionSms{
		smsType: SmsTypeCommon,
		mobiles: []string{target},
	}
}

// NotifySms 配置通知短信
func NotifySms(target string) *optionSms {
	return &optionSms{
		smsType: SmsTypeNotify,
		mobiles: []string{target},
	}
}

// AdvertisingSms 配置营销（广告）短信
func AdvertisingSms(targets ...string) *optionSms {
	return &optionSms{
		smsType: SmsTypeAdvertising,
		mobiles: targets,
	}
}

func (s *optionSms) apply(options *options) {
	options.chuangcacheSms.smsType = s.smsType
	options.chuangcacheSms.mobiles = s.mobiles
}
