package una

import (
	`fmt`
)

var _ option = (*optionEmailContent)(nil)

type optionEmailContent struct {
	// 发件人
	from string `validate:"required"`
	// 内容
	content string `validate:"required"`
	// 邮件类型
	emailType EmailType `validate:"required,oneof=html plain"`
	// 发送地址列表
	to []string `validate:"required,dive,emailConfig"`
}

// HtmlEmail 配置富文本邮件
func HtmlEmail(name string, from string, to ...string) *optionEmailContent {
	return &optionEmailContent{
		from:      fmt.Sprintf("%s <%s>", name, from),
		emailType: EmailTypeHtml,
		to:        to,
	}
}

// PlainEmail 配置普通邮件
func PlainEmail(name string, from string, to ...string) *optionEmailContent {
	return &optionEmailContent{
		from:      fmt.Sprintf("%s <%s>", name, from),
		emailType: EmailTypePlain,
		to:        to,
	}
}

func (ec *optionEmailContent) apply(options *options) {
	options.email.from = ec.from
	options.email.emailType = ec.emailType
	options.email.to = ec.to
}
