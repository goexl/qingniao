package qingniao

import (
	"github.com/goexl/http"
)

var (
	_ option = (*optionHttp)(nil)
	_        = HttpClient
)

type optionHttp struct {
	client *http.Client
}

// HttpClient 配置客户端
func HttpClient(client *http.Client) *optionHttp {
	return &optionHttp{
		client: client,
	}
}

func (h *optionHttp) apply(options *options) {
	options.http = h.client
}
