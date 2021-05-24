package una

import (
	`runtime`
)

type options struct {
	// 短信配置
	chuangcache chuangcacheConfig
	// 邮件配置
	email emailConfig
	// 池数量
	poolSize int
}

func defaultOptions() *options {
	return &options{
		chuangcache: chuangcacheConfig{
			smsType: SmsTypeCommon,
		},
		email: emailConfig{
			port:      465,
			emailType: EmailTypeHtml,
		},
		poolSize: runtime.NumCPU() + 1,
	}
}
