package una

var _ option = (*optionEmailAuth)(nil)

type optionEmailAuth struct {
	// 邮箱地址
	host string `validate:"required"`
	// 端口
	port int `validate:"required"`
	// 用户名
	username string `validate:"required"`
	// 密码
	password string `validate:"required"`
}

// EmailAuth 配置邮件服务
func EmailAuth(host string, port int, username string, password string) *optionEmailAuth {
	return &optionEmailAuth{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}
}

func (a *optionEmailAuth) apply(options *options) {
	options.email.host = a.host
	options.email.port = a.port
	options.email.username = a.username
	options.email.password = a.password
	options.unaType = TypeEmail
}
