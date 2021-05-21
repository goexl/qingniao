package una

type options struct {
	// 短信配置
	sms sms
	// 邮件配置
	email email
	// 创世云授权配置
	akSk akSk
	// 主题，邮件使用
	subject string
}

func defaultOptions() *options {
	return &options{}
}
