package qingniao

var _ option = (*optionTemplate)(nil)

type optionTemplate struct {
	data interface{}
}

// Template 配置主题
func Template(data interface{}) *optionTemplate {
	return &optionTemplate{
		data: data,
	}
}

func (t *optionTemplate) apply(options *options) {
	options.template = true
	options.data = t.data
}
