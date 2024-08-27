package builder

import (
	"github.com/goexl/qingniao/internal/internal"
	"github.com/goexl/qingniao/internal/internal/constant"
	"github.com/goexl/qingniao/internal/internal/executor"
	"github.com/goexl/qingniao/internal/param"
)

type Wechat struct {
	params *param.Sender
	sender *Sender
	chain  internal.Wechat
}

func NewWechat(params *param.Sender, sender *Sender) *Wechat {
	return &Wechat{
		params: params,
		sender: sender,
	}
}

func (w *Wechat) ServerChan(key string) (wechat *Wechat) {
	w.chain = executor.NewServerChan(key, w.params.Http, w.params.Logger)
	wechat = w

	return
}

func (w *Wechat) Build() (sender *Sender) {
	w.sender.wechat[constant.ExecutorServerChain] = w.chain
	sender = w.sender

	return
}
