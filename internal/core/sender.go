package core

import (
	"github.com/goexl/qingniao/internal/internal/deliver"
)

type Sender interface {
	// Wechat 发送微信
	Wechat(title string, content string) *deliver.Wechat

	// Email 发送邮件
	Email(address string, subject string, content string) *deliver.Email

	// Sms 发送短信
	Sms(mobile string, content string) *deliver.Sms
}
