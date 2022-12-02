package qingniao

import (
	"time"

	"github.com/goexl/xiren"
	"github.com/jordan-wright/email"
)

type emailDeliver struct {
	typ     emailType `validate:"required,oneof=html plain"`
	subject string    `validate:"required"`
	content string    `validate:"required"`
	from    string    `validate:"required"`
	to      []string  `validate:"required,dive,email"`
	cc      []string  `validate:"omitempty,dive,email"`
	bcc     []string  `validate:"omitempty,dive,email"`
	timeout time.Duration

	pool *email.Pool
}

func newEmailDeliver(from string, subject string, content string, pool *email.Pool) *emailDeliver {
	return &emailDeliver{
		subject: subject,
		content: content,
		from:    from,
		pool:    pool,
		timeout: 10 * time.Second,
	}
}

func (ed *emailDeliver) Send() (err error) {
	if err = xiren.Struct(ed); nil != err {
		return
	}

	em := email.NewEmail()
	em.From = ed.from
	em.To = ed.to
	em.Bcc = ed.bcc
	em.Cc = ed.cc
	em.Subject = ed.subject
	switch ed.typ {
	case emailTypeHtml:
		em.HTML = []byte(ed.content)
	case emailTypePlain:
		em.Text = []byte(ed.content)
	default:
		em.HTML = []byte(ed.content)
	}
	err = ed.pool.Send(em, ed.timeout)

	return
}

func (ed *emailDeliver) From(from string) *emailDeliver {
	ed.from = from

	return ed
}

func (ed *emailDeliver) To(to string) *emailDeliver {
	ed.to = append(ed.to, to)

	return ed
}

func (ed *emailDeliver) Cc(cc string) *emailDeliver {
	ed.cc = append(ed.cc, cc)

	return ed
}

func (ed *emailDeliver) Bcc(bcc string) *emailDeliver {
	ed.bcc = append(ed.bcc, bcc)

	return ed
}

func (ed *emailDeliver) Subject(subject string) *emailDeliver {
	ed.subject = subject

	return ed
}

func (ed *emailDeliver) Html() *emailDeliver {
	ed.typ = emailTypeHtml

	return ed
}

func (ed *emailDeliver) Plain() *emailDeliver {
	ed.typ = emailTypePlain

	return ed
}

func (ed *emailDeliver) Timeout(timeout time.Duration) *emailDeliver {
	ed.timeout = timeout

	return ed
}
