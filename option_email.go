package una

var _ option = (*optionEmailAuth)(nil)

type optionEmailAuth struct {
	email email
}

// HtmlEmail 配置富文本邮件
func HtmlEmail(host string, username string, password string) *optionEmailAuth {
	return &optionEmailAuth{
		email: email{
			host:      host,
			username:  username,
			password:  password,
			emailType: EmailTypeHtml,
		},
	}
}

// PlainEmail 配置普通文本邮件
func PlainEmail(host string, username string, password string) *optionEmailAuth {
	return &optionEmailAuth{
		email: email{
			host:      host,
			username:  username,
			password:  password,
			emailType: EmailTypePlain,
		},
	}
}

func (as *optionEmailAuth) apply(options *options) {
	options.email = as.email
}
