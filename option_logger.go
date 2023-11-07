package qingniao

import (
	"github.com/goexl/log"
)

var (
	_ option = (*optionLogger)(nil)
	_        = Logger
)

type optionLogger struct {
	logger log.Logger
}

// Logger 配置日志
func Logger(logger log.Logger) *optionLogger {
	return &optionLogger{
		logger: logger,
	}
}

func (l *optionLogger) apply(options *options) {
	options.logger = l.logger
}
