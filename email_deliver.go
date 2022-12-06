package qingniao

import (
	"context"
	"time"
)

type emailDeliver struct {
	Type          emailType `validate:"required,oneof=html plain"`
	// 避免和方法冲突不得已命名
	SubjectString string    `validate:"required"`
	Content       string    `validate:"required"`
	// 避免和方法冲突不得已命名
	FromAddress  string    `validate:"required"`
	// 避免和方法冲突不得已命名
	ToSlice      []string  `validate:"required,dive,email"`
	// 避免和方法冲突不得已命名
	CcSlice      []string  `validate:"omitempty,dive,email"`
	// 避免和方法冲突不得已命名
	BccSlice     []string  `validate:"omitempty,dive,email"`
	timeout      time.Duration

	executor emailExecutor
}

func newEmailDeliver(addresses []string, subject string, content string, executor emailExecutor) *emailDeliver {
	return &emailDeliver{
		SubjectString: subject,
		Content:       content,
		ToSlice:       addresses,
		executor:      executor,
		timeout:       10 * time.Second,
	}
}

func (ed *emailDeliver) Send(ctx context.Context) (string, error) {
	return ed.executor.send(ctx, ed)
}

func (ed *emailDeliver) From(from string) *emailDeliver {
	ed.FromAddress = from

	return ed
}

func (ed *emailDeliver) To(to ...string) *emailDeliver {
	ed.ToSlice = append(ed.ToSlice, to...)

	return ed
}

func (ed *emailDeliver) Cc(cc ...string) *emailDeliver {
	ed.CcSlice = append(ed.CcSlice, cc...)

	return ed
}

func (ed *emailDeliver) Bcc(bcc ...string) *emailDeliver {
	ed.BccSlice = append(ed.BccSlice, bcc...)

	return ed
}

func (ed *emailDeliver) Subject(subject string) *emailDeliver {
	ed.SubjectString = subject

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
