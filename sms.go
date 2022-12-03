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

func (s *Sms) To(mobile string, content string) *smsDeliver {
	return newSmsDeliver(mobile, content, s.executor)
}
