package una

import (
	`fmt`
)

type emailConfig struct {
	// 邮箱地址
	host string `validate:"required"`
	// 端口
	port int `validate:"required"`
	// 用户名
	username string `validate:"required"`
	// 密码
	password string `validate:"required"`
	// 邮件类型
	emailType EmailType `validate:"required,oneof=html plain"`
	// 邮件主题
	subject string `validate:"required"`
	// 发送人
	from string `validate:"required"`
	// 发送地址列表
	to []string `validate:"required,dive,email"`
	// 抄送列表
	cc []string `validate:"omitempty,dive,email"`
	// 秘送列表
	bcc []string `validate:"omitempty,dive,email"`
}

func (ec *emailConfig) key() string {
	return fmt.Sprintf("%s-%s", ec.host, ec.username)
}
