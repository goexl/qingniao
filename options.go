package una

type options struct {
	// 过期时间
	smsType SmsType
}

func defaultOptions() *options {
	return &options{
		smsType: SmsTypeCommon,
	}
}
