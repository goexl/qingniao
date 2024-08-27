package builder

import (
	"github.com/goexl/qingniao/internal/internal"
	"github.com/goexl/qingniao/internal/internal/executor"
	"github.com/goexl/qingniao/internal/param"
)

type Wechat struct {
	*base[Wechat]

	params *param.Sender
	sender *Sender
	chain  internal.Wechat
}

func NewWechat(params *param.Sender, sender *Sender) (wechat *Wechat) {
	wechat = new(Wechat)
	wechat.base = newBase(wechat)

	wechat.params = params
	wechat.sender = sender

	return
}

func (w *Wechat) Serverchan(key string) (wechat *Wechat) {
	w.chain = executor.NewServerchan(key, w.params.Http, w.params.Logger)
	wechat = w

	return
}

func (w *Wechat) Build() (sender *Sender) {
	w.sender.wechat[w.base.params.Label] = w.chain
	sender = w.sender

	return
}
