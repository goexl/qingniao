package qingniao

import (
	"github.com/goexl/http"
	"github.com/goexl/log"
)

type smsBuilder struct {
	http   *http.Client
	logger log.Logger
}

func newSmsBuilder(http *http.Client, logger log.Logger) *smsBuilder {
	return &smsBuilder{
		http:   http,
		logger: logger,
	}
}

func (sb *smsBuilder) Chuangcache(ak string, sk string) *Sms {
	return newSms(newChuangcache(ak, sk, sb.http, sb.logger))
}
