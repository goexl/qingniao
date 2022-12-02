package qingniao

import (
	"github.com/go-resty/resty/v2"
)

var _ = New

// Sender 发送者
type Sender struct {
	http *resty.Client
}

// New 创建发送者
func New(client *resty.Client) *Sender {
	return &Sender{
		http: client,
	}
}

func (s *Sender) Email(host string, port int) *emailBuilder {
	return newEmailBuilder(host, port)
}

func (s *Sender) Sms() *smsBuilder {
	return newSmsBuilder()
}
