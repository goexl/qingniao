package qingniao

import (
	"github.com/go-resty/resty/v2"
	"github.com/goexl/simaqian"
)

var _ = New

// Sender 发送者
type Sender struct {
	http   *resty.Client
	logger simaqian.Logger
}

// New 创建发送者
func New(http *resty.Client, logger simaqian.Logger) *Sender {
	return &Sender{
		http:   http,
		logger: logger,
	}
}

func (s *Sender) Email() *emailBuilder {
	return newEmailBuilder(s.http, s.logger)
}

func (s *Sender) Sms() *smsBuilder {
	return newSmsBuilder(s.http, s.logger)
}
