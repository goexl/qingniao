package builder

import (
	"github.com/goexl/qingniao/internal/internal"
	"github.com/goexl/qingniao/internal/param"
)

type Email struct {
	*base[Email]

	params *param.Sender
	sender *Sender
	direct internal.Email
}

func NewEmail(params *param.Sender, sender *Sender) (email *Email) {
	email = new(Email)
	email.base = newBase(email)

	email.params = params
	email.sender = sender

	return
}

func (e *Email) Smtp(host string, port int) *Smtp {
	return NewSmtp(host, port, e.params.Logger, e)
}

func (e *Email) Build() (sender *Sender) {
	e.sender.email[e.base.params.Label] = e.direct
	sender = e.sender

	return
}
