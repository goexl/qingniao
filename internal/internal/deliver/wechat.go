package deliver

import (
	"context"

	"github.com/goexl/qingniao/internal/internal"
	"github.com/goexl/qingniao/internal/internal/constant"
	"github.com/goexl/qingniao/internal/internal/executor/deliver"
	"github.com/goexl/qingniao/internal/kernel"
	"github.com/goexl/xiren"
)

type Wechat struct {
	title   string
	content string

	picker  *picker[internal.Wechat]
	current constant.Executor
}

func NewWechat(title string, content string, executors map[constant.Executor]internal.Wechat) *Wechat {
	return &Wechat{
		title:   title,
		content: content,

		picker: newPicker(executors),
	}
}

func (w *Wechat) Send(ctx context.Context) (id string, status kernel.Status, err error) {
	message := new(deliver.Wechat)
	message.Title = w.title
	message.Content = w.content
	if se := xiren.Struct(message); nil != se {
		err = se
	} else if executor, pe := w.picker.pick(w.current, "微信"); nil != pe {
		err = pe
	} else {
		id, status, err = executor.Send(ctx, message)
	}

	return
}

func (w *Wechat) ServerChain() (wechat *Wechat) {
	w.current = constant.ExecutorServerChain
	wechat = w

	return
}
