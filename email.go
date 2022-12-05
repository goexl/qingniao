package qingniao

// Email 邮件
type Email struct {
	executor emailExecutor
}

func newEmail(executor emailExecutor) *Email {
	return &Email{
		executor: executor,
	}
}

func (e *Email) Deliver(subject string, content string, addresses ...string) *emailDeliver {
	return newEmailDeliver(addresses, subject, content, e.executor)
}
