package qingniao

import (
	"github.com/go-resty/resty/v2"
	"github.com/goexl/simaqian"
)

type emailBuilder struct {
	http   *resty.Client
	logger simaqian.Logger
}

func newEmailBuilder(http *resty.Client, logger simaqian.Logger) *emailBuilder {
	return &emailBuilder{
		http:   http,
		logger: logger,
	}
}

func (eb *emailBuilder) Direct(host string, port int) *directBuilder {
	return newDirectBuilder(host, port, eb.logger)
}
