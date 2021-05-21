package una

import (
	`context`
	`net/smtp`
	`strings`

	`github.com/storezhang/validatorx`
)

type (
	// Email 邮件通知
	Email struct{}

	email struct {
		// 邮箱地址
		host string `validate:"required"`
		// 用户名
		username string `validate:"required"`
		// 密码
		password string `validate:"required"`
		// 邮件类型
		emailType EmailType `validate:"required,oneof=html plain"`
	}
)

// NewEmail 创建普通邮件
func NewEmail() *Email {
	return &Email{}
}

func (e *Email) Send(_ context.Context, content string, opts ...option) (id string, err error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}
	if err = validatorx.Validate(options.email); nil != err {
		return
	}

	auth := smtp.PlainAuth("", options.email.username, options.email.password, options.email.host)
	var contentType string
	switch options.email.emailType {
	case EmailTypePlain:
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	case EmailTypeHtml:
		contentType = "Content-Type: text/html" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + sendUserName + "<" + user + ">" + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err = smtp.SendMail(options.email.host, auth, user, send_to, msg)

	return
}
