package qingniao

import (
	"github.com/goexl/http"
	"github.com/goexl/log"
)

var _ = NewOptions

type (
	option interface {
		apply(options *options)
	}

	options struct {
		http   *http.Client
		logger log.Logger
	}
)

// NewOptions 方便选项书写
func NewOptions(opts ...option) []option {
	return opts
}

func defaultOptions() *options {
	return &options{}
}
