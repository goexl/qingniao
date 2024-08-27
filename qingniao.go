package qingniao

import (
	"github.com/goexl/qingniao/internal/builder"
)

// New 创建发送器
func New() *builder.Sender {
	return builder.NewSender()
}
