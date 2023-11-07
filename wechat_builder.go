package qingniao

import (
	"github.com/goexl/http"
	"github.com/goexl/log"
)

type wechatBuilder struct {
	http   *http.Client
	logger log.Logger
}

func newWechatBuilder(http *http.Client, logger log.Logger) *wechatBuilder {
	return &wechatBuilder{
		http:   http,
		logger: logger,
	}
}

func (wb *wechatBuilder) ServerChan(key string) *Wechat {
	return newWechat(newServerChan(key, wb.http, wb.logger))
}
