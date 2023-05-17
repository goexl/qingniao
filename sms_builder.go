package qingniao

import (
	"github.com/go-resty/resty/v2"
	"github.com/goexl/simaqian"
)

type smsBuilder struct {
	http   *resty.Client
	logger simaqian.Logger
}

func newSmsBuilder(http *resty.Client, logger simaqian.Logger) *smsBuilder {
	return &smsBuilder{
		http:   http,
		logger: logger,
	}
}

func (sb *smsBuilder) Chuangcache(ak string, sk string) *Sms {
	return newSms(newChuangcache(ak, sk, sb.http, sb.logger))
}
