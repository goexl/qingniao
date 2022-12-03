package qingniao

import (
	"github.com/go-resty/resty/v2"
)

var (
	_ option = (*optionHttp)(nil)
	_        = HttpClient
)

type optionHttp struct {
	client *resty.Client
}

// HttpClient 配置客户端
func HttpClient(client *resty.Client) *optionHttp {
	return &optionHttp{
		client: client,
	}
}

func (h *optionHttp) apply(options *options) {
	options.http = h.client
}
