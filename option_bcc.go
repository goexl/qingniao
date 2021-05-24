package una

var _ option = (*optionBCC)(nil)

type optionBCC struct {
	bcc []string
}

// BCC 配置秘送列表
func BCC(emails ...string) *optionBCC {
	return &optionBCC{
		bcc: emails,
	}
}

func (as *optionBCC) apply(options *options) {
	options.mail.bcc = as.bcc
}
