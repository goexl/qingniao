package builder

import (
	"github.com/goexl/qingniao/internal/internal"
	"github.com/goexl/qingniao/internal/internal/executor"
	"github.com/goexl/qingniao/internal/param"
)

type Sms struct {
	*base[Sms]

	params *param.Sender
	sender *Sender
	cache  internal.Sms
}

func NewSms(params *param.Sender, sender *Sender) (sms *Sms) {
	sms = new(Sms)
	sms.base = newBase(sms)

	sms.params = params
	sms.sender = sender

	return
}

func (s *Sms) Chuangcache(ak string, sk string) (sms *Sms) {
	s.cache = executor.NewChuangcache(ak, sk, s.params.Http, s.params.Logger)
	sms = s

	return
}

func (s *Sms) Build() (sender *Sender) {
	s.sender.sms[s.base.params.Label] = s.cache
	sender = s.sender

	return
}
