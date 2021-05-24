package una

var _ option = (*optionCC)(nil)

type optionCC struct {
	cc []string
}

// CC 配置抄送列表
func CC(emails ...string) *optionCC {
	return &optionCC{
		cc: emails,
	}
}

func (as *optionCC) apply(options *options) {
	options.email.cc = as.cc
}
