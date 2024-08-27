package deliver

import (
	"context"
	"time"

	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	"github.com/goexl/qingniao/internal/internal"
	"github.com/goexl/qingniao/internal/internal/constant"
	"github.com/goexl/qingniao/internal/internal/executor/deliver"
	"github.com/goexl/xiren"
)

type Email struct {
	base

	et      constant.EmailType
	subject string
	content string
	from    string
	to      []string
	cc      []string
	bcc     []string
	timeout time.Duration

	executors map[string]internal.Email
}

func NewEmail(address string, subject string, content string, executors map[string]internal.Email) (email *Email) {
	return &Email{
		base: newBase(),

		subject: subject,
		content: content,
		to:      []string{address},
		timeout: 10 * time.Second,

		executors: executors,
	}
}

func (e *Email) Send(ctx context.Context) (id string, err error) {
	message := new(deliver.Email)
	message.Type = e.et
	message.Subject = e.subject
	message.Content = e.content
	message.From = e.from
	message.To = e.to
	message.Cc = e.cc
	message.Bcc = e.bcc
	message.Timeout = e.timeout

	label := e.base.params.Label
	if se := xiren.Struct(message); nil != se {
		err = se
	} else if executor, ok := e.executors[label]; !ok {
		err = exception.New().Message("没有找到邮件执行器").
			Field(field.New("executors", e.executors)).Field(field.New("label", label)).
			Build()
	} else {
		id, err = executor.Send(ctx, message)
	}

	return
}

func (e *Email) From(from string) (email *Email) {
	e.from = from
	email = e

	return
}

func (e *Email) To(to string, extras ...string) (email *Email) {
	e.to = append(e.to, to)
	e.to = append(e.to, extras...)
	email = e

	return
}

func (e *Email) Cc(cc string, extras ...string) (email *Email) {
	e.cc = append(e.cc, cc)
	e.cc = append(e.cc, extras...)
	email = e

	return
}

func (e *Email) Bcc(bcc string, extras ...string) (email *Email) {
	e.bcc = append(e.bcc, bcc)
	e.bcc = append(e.bcc, extras...)
	email = e

	return
}

func (e *Email) Subject(subject string) (email *Email) {
	e.subject = subject
	email = e

	return
}

func (e *Email) Html() (email *Email) {
	e.et = constant.EmailTypeHtml
	email = e

	return
}

func (e *Email) Plain() (email *Email) {
	e.et = constant.EmailTypePlain
	email = e

	return
}

func (e *Email) Timeout(timeout time.Duration) (email *Email) {
	e.timeout = timeout
	email = e

	return
}
