package qingniao

import (
	"github.com/goexl/http"
	"github.com/goexl/log"
)

type emailBuilder struct {
	http   *http.Client
	logger log.Logger
}

func newEmailBuilder(http *http.Client, logger log.Logger) *emailBuilder {
	return &emailBuilder{
		http:   http,
		logger: logger,
	}
}

func (eb *emailBuilder) Direct(host string, port int) *directBuilder {
	return newDirectBuilder(host, port, eb.logger)
}
