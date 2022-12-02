package qingniao

import (
	"runtime"
)

type options struct {
	// 短信配置
	chuangcache chuangcacheConfig
	// 邮件配置
	email emailDeliver
	// 池数量
	poolSize int
	// 是否是模板
	template bool
	// 模板数据
	data interface{}

	// 类型
	unaType Type
}

func defaultOptions() *options {
	return &options{
		chuangcache: chuangcacheConfig{
			smsType: SmsTypeCommon,
		},
		email: emailDeliver{
			port: 465,
			typ:  emailTypeHtml,
		},
		poolSize: runtime.NumCPU() + 1,
		template: false,
		unaType:  TypeEmail,
	}
}
