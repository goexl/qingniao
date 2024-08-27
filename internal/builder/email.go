package builder

import (
	"github.com/goexl/qingniao/internal/internal"
	"github.com/goexl/qingniao/internal/internal/constant"
	"github.com/goexl/qingniao/internal/param"
)

type Email struct {
	params *param.Sender
	sender *Sender
	direct internal.Email
}

func NewEmail(params *param.Sender, sender *Sender) *Email {
	return &Email{
		params: params,
		sender: sender,
	}
}

func (e *Email) Direct(host string, port int) *Direct {
	return NewDirect(host, port, e.params.Logger, e)
}

func (e *Email) Build() (sender *Sender) {
	e.sender.email[constant.ExecutorDirect] = e.direct
	sender = e.sender

	return
}
