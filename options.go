package una

import (
	`runtime`
)

type options struct {
	// 发件人
	from string
	// 收件人
	to []string
	// 短信配置
	sms sms
	// 邮件配置
	mail mail
	// 创世云授权配置
	akSk akSk
	// 池数量
	poolSize int
}

func defaultOptions() *options {
	return &options{
		sms: sms{
			smsType: SmsTypeCommon,
		},
		mail: mail{
			port:      465,
			emailType: EmailTypeHtml,
		},
		poolSize: runtime.NumCPU() + 1,
	}
}
