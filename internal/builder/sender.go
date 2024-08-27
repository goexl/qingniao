package builder

import (
	"github.com/goexl/http"
	"github.com/goexl/log"
	"github.com/goexl/qingniao/internal/core"
	"github.com/goexl/qingniao/internal/internal"
	"github.com/goexl/qingniao/internal/internal/kernel"
	"github.com/goexl/qingniao/internal/param"
)

type Sender struct {
	params *param.Sender
	email  map[string]internal.Email
	sms    map[string]internal.Sms
	wechat map[string]internal.Wechat
}

func NewSender() *Sender {
	return &Sender{
		params: param.NewSender(),
		email:  make(map[string]internal.Email),
		sms:    make(map[string]internal.Sms),
		wechat: make(map[string]internal.Wechat),
	}
}

func (s *Sender) Logger(logger log.Logger) (sender *Sender) {
	s.params.Logger = logger
	sender = s

	return
}

func (s *Sender) Http(http *http.Client) (sender *Sender) {
	s.params.Http = http
	sender = s

	return
}

func (s *Sender) Email() *Email {
	return NewEmail(s.params, s)
}

func (s *Sender) Sms() *Sms {
	return NewSms(s.params, s)
}

func (s *Sender) Wechat() *Wechat {
	return NewWechat(s.params, s)
}

func (s *Sender) Build() core.Sender {
	return kernel.NewSender(s.email, s.sms, s.wechat)
}
