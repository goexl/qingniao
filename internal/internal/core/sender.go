package core

import (
	"github.com/goexl/qingniao/internal/core"
	"github.com/goexl/qingniao/internal/internal"
	"github.com/goexl/qingniao/internal/internal/deliver"
)

var _ core.Sender = (*Sender)(nil)

type Sender struct {
	email  map[string]internal.Email
	sms    map[string]internal.Sms
	wechat map[string]internal.Wechat
}

func NewSender(
	email map[string]internal.Email,
	sms map[string]internal.Sms,
	wechat map[string]internal.Wechat,
) *Sender {
	return &Sender{
		email:  email,
		sms:    sms,
		wechat: wechat,
	}
}

func (s Sender) Wechat(title string, content string) *deliver.Wechat {
	return deliver.NewWechat(title, content, s.wechat)
}

func (s Sender) Email(address string, subject string, content string) *deliver.Email {
	return deliver.NewEmail(address, subject, content, s.email)
}

func (s Sender) Sms(mobile string, content string) *deliver.Sms {
	return deliver.NewSms(mobile, content, s.sms)
}
