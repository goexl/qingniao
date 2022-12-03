package qingniao

import (
	"github.com/go-resty/resty/v2"
	"github.com/goexl/simaqian"
)

var _ = New

// Sender 发送者
type Sender struct {
	http   *resty.Client
	logger simaqian.Logger
}

// New 创建发送者
func New(opts ...option) *Sender {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}
	if nil == _options.http {
		_options.http = resty.New()
	}
	if nil == _options.logger {
		_options.logger = simaqian.Must()
	}

	return &Sender{
		http:   _options.http,
		logger: _options.logger,
	}
}

func (s *Sender) Email() *emailBuilder {
	return newEmailBuilder(s.http, s.logger)
}

func (s *Sender) Sms() *smsBuilder {
	return newSmsBuilder(s.http, s.logger)
}
