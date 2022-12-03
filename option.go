package qingniao

import (
	"github.com/go-resty/resty/v2"
	"github.com/goexl/simaqian"
)

var _ = NewOptions

type (
	option interface {
		apply(options *options)
	}

	options struct {
		http   *resty.Client
		logger simaqian.Logger
	}
)

// NewOptions 方便选项书写
func NewOptions(opts ...option) []option {
	return opts
}

func defaultOptions() *options {
	return &options{}
}
