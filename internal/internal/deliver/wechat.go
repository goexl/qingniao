package deliver

import (
	"context"

	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	"github.com/goexl/qingniao/internal/internal"
	"github.com/goexl/qingniao/internal/internal/executor/deliver"
	"github.com/goexl/qingniao/internal/kernel"
	"github.com/goexl/xiren"
)

type Wechat struct {
	*base[Wechat]

	title   string
	content string

	executors map[string]internal.Wechat
}

func NewWechat(title string, content string, executors map[string]internal.Wechat) (wechat *Wechat) {
	wechat = new(Wechat)
	wechat.base = newBase(wechat)

	wechat.title = title
	wechat.content = content

	wechat.executors = executors

	return
}

func (w *Wechat) Send(ctx context.Context) (id string, status kernel.Status, err error) {
	message := new(deliver.Wechat)
	message.Title = w.title
	message.Content = w.content

	label := w.base.params.Label
	if se := xiren.Struct(message); nil != se {
		err = se
	} else if executor, ok := w.executors[label]; !ok {
		err = exception.New().Message("没有找到微信执行器").
			Field(field.New("executors", w.executors)).Field(field.New("label", label)).
			Build()
	} else {
		id, status, err = executor.Send(ctx, message)
	}

	return
}
