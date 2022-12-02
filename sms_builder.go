package qingniao

import (
	"github.com/go-resty/resty/v2"
)

type smsBuilder struct {
	http     *resty.Client
	executor smsExecutor
}

func newSmsBuilder(http *resty.Client) *smsBuilder {
	return &smsBuilder{
		http: http,
	}
}

func (sb *smsBuilder) Chuangcache(ak string, sk string) *smsBuilder {
	sb.executor = newChuangcache(sb.http)

	return sb
}
