package builder

import (
	"github.com/goexl/qingniao/internal/internal"
	"github.com/goexl/qingniao/internal/param"
)

type Email struct {
	base

	params *param.Sender
	sender *Sender
	direct internal.Email
}

func NewEmail(params *param.Sender, sender *Sender) *Email {
	return &Email{
		base: newBase(),

		params: params,
		sender: sender,
	}
}

func (e *Email) Smtp(host string, port int) *Smtp {
	return NewSmtp(host, port, e.params.Logger, e)
}

func (e *Email) Build() (sender *Sender) {
	e.sender.email[e.base.params.Label] = e.direct
	sender = e.sender

	return
}
