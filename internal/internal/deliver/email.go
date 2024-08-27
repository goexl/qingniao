package deliver

import (
	"context"
	"time"

	"github.com/goexl/qingniao/internal/internal"
	"github.com/goexl/qingniao/internal/internal/constant"
	"github.com/goexl/qingniao/internal/internal/executor/deliver"
	"github.com/goexl/xiren"
)

type Email struct {
	et      constant.EmailType
	subject string
	content string
	from    string
	to      []string
	cc      []string
	bcc     []string
	timeout time.Duration

	picker  *picker[internal.Email]
	current constant.Executor
}

func NewEmail(address string, subject string, content string, executors map[constant.Executor]internal.Email) (email *Email) {
	return &Email{
		subject: subject,
		content: content,
		to:      []string{address},
		timeout: 10 * time.Second,

		picker: newPicker(executors),
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
	if se := xiren.Struct(message); nil != se {
		err = se
	} else if executor, pe := e.picker.pick(e.current, "邮件"); nil != pe {
		err = pe
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

func (e *Email) Direct() (email *Email) {
	e.current = constant.ExecutorDirect
	email = e

	return
}
