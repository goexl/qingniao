package qingniao

import (
	"context"
	"time"
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

	executor emailExecutor
}

func newEmailDeliver(addresses []string, subject string, content string, executor emailExecutor) *emailDeliver {
	return &emailDeliver{
		subject:  subject,
		content:  content,
		to:       addresses,
		executor: executor,
		timeout:  10 * time.Second,
	}
}

func (ed *emailDeliver) Send(ctx context.Context) (string, error) {
	return ed.executor.send(ctx, ed)
}

func (ed *emailDeliver) From(from string) *emailDeliver {
	ed.from = from

	return ed
}

func (ed *emailDeliver) To(to ...string) *emailDeliver {
	ed.to = append(ed.to, to...)

	return ed
}

func (ed *emailDeliver) Cc(cc ...string) *emailDeliver {
	ed.cc = append(ed.cc, cc...)

	return ed
}

func (ed *emailDeliver) Bcc(bcc ...string) *emailDeliver {
	ed.bcc = append(ed.bcc, bcc...)

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
