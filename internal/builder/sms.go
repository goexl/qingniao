package builder

import (
	"github.com/goexl/qingniao/internal/internal"
	"github.com/goexl/qingniao/internal/internal/constant"
	"github.com/goexl/qingniao/internal/internal/executor"
	"github.com/goexl/qingniao/internal/param"
)

type Sms struct {
	params *param.Sender
	sender *Sender
	cache  internal.Sms
}

func NewSms(params *param.Sender, sender *Sender) *Sms {
	return &Sms{
		params: params,
		sender: sender,
	}
}

func (s *Sms) Chuangcache(ak string, sk string) (sms *Sms) {
	s.cache = executor.NewChuangcache(ak, sk, s.params.Http, s.params.Logger)
	sms = s

	return
}

func (s *Sms) Build() (sender *Sender) {
	s.sender.sms[constant.ExecutorChuangcache] = s.cache
	sender = s.sender

	return
}
