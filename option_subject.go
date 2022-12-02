package qingniao

var _ option = (*optionSubject)(nil)

type optionSubject struct {
	subject string
}

// Subject 配置主题
func Subject(subject string) *optionSubject {
	return &optionSubject{
		subject: subject,
	}
}

func (s *optionSubject) apply(options *options) {
	options.email.subject = s.subject
}
