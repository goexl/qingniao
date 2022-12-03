package qingniao

import (
	"github.com/goexl/simaqian"
)

var (
	_ option = (*optionLogger)(nil)
	_        = Logger
)

type optionLogger struct {
	logger simaqian.Logger
}

// Logger 配置日志
func Logger(logger simaqian.Logger) *optionLogger {
	return &optionLogger{
		logger: logger,
	}
}

func (l *optionLogger) apply(options *options) {
	options.logger = l.logger
}
