package qingniao

import (
	"context"
	"time"
)

type emailDeliver struct {
	Type         emailType `validate:"required,oneof=html plain"`
	SubjectField string    `validate:"required"`
	Content      string    `validate:"required"`
	FromField    string    `validate:"required"`
	ToField      []string  `validate:"required,dive,email"`
	CcField      []string  `validate:"omitempty,dive,email"`
	BccField     []string  `validate:"omitempty,dive,email"`
	timeout      time.Duration

	executor emailExecutor
}

func newEmailDeliver(addresses []string, subject string, content string, executor emailExecutor) *emailDeliver {
	return &emailDeliver{
		SubjectField: subject,
		Content:      content,
		ToField:      addresses,
		executor:     executor,
		timeout:      10 * time.Second,
	}
}

func (ed *emailDeliver) Send(ctx context.Context) (string, error) {
	return ed.executor.send(ctx, ed)
}

func (ed *emailDeliver) From(from string) *emailDeliver {
	ed.FromField = from

	return ed
}

func (ed *emailDeliver) To(to ...string) *emailDeliver {
	ed.ToField = append(ed.ToField, to...)

	return ed
}

func (ed *emailDeliver) Cc(cc ...string) *emailDeliver {
	ed.CcField = append(ed.CcField, cc...)

	return ed
}

func (ed *emailDeliver) Bcc(bcc ...string) *emailDeliver {
	ed.BccField = append(ed.BccField, bcc...)

	return ed
}

func (ed *emailDeliver) Subject(subject string) *emailDeliver {
	ed.SubjectField = subject

	return ed
}

func (ed *emailDeliver) Html() *emailDeliver {
	ed.Type = emailTypeHtml

	return ed
}

func (ed *emailDeliver) Plain() *emailDeliver {
	ed.Type = emailTypePlain

	return ed
}

func (ed *emailDeliver) Timeout(timeout time.Duration) *emailDeliver {
	ed.timeout = timeout

	return ed
}
