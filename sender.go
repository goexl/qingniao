package qingniao

import (
	"github.com/goexl/http"
	"github.com/goexl/log"
)

var _ = New

// Sender 发送者
type Sender struct {
	http   *http.Client
	logger log.Logger
}

// New 创建发送者
func New(opts ...option) *Sender {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}
	if nil == _options.http {
		_options.http = http.New().Build()
	}
	if nil == _options.logger {
		_options.logger = log.New().Apply()
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

func (s *Sender) Wechat() *wechatBuilder {
	return newWechatBuilder(s.http, s.logger)
}
