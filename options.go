package una

import (
	`runtime`
)

type options struct {
	// 短信配置
	chuangcacheSms chuangcacheSmsConfig
	// 邮件配置
	email emailConfig
	// 池数量
	poolSize int
}

func defaultOptions() *options {
	return &options{
		chuangcacheSms: chuangcacheSmsConfig{
			smsType: SmsTypeCommon,
		},
		email: emailConfig{
			port:      465,
			emailType: EmailTypeHtml,
		},
		poolSize: runtime.NumCPU() + 1,
	}
}
