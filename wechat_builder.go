package qingniao

import (
	"github.com/go-resty/resty/v2"
	"github.com/goexl/simaqian"
)

type wechatBuilder struct {
	http   *resty.Client
	logger simaqian.Logger
}

func newWechatBuilder(http *resty.Client, logger simaqian.Logger) *wechatBuilder {
	return &wechatBuilder{
		http:   http,
		logger: logger,
	}
}

func (wb *wechatBuilder) ServerChan(key string) *Wechat {
	return newWechat(newServerChan(key, wb.http, wb.logger))
}
