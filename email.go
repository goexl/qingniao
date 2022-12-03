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

func (e *Email) To(addr string, subject string, content string) *emailDeliver {
	return newEmailDeliver(addr, subject, content, e.executor)
}
