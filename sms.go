package qingniao

// Sms 短信
type Sms struct {
	executor smsExecutor
}

func newSms(executor smsExecutor) *Sms {
	return &Sms{
		executor: executor,
	}
}

func (s *Sms) To(template string, mobile string, content string) *smsDeliver {
	return newSmsDeliver(template, mobile, content, s.executor)
}
