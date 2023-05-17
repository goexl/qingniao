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

func (s *Sms) Deliver(content string, mobiles ...string) *smsDeliver {
	return newSmsDeliver(mobiles, content, s.executor)
}
