package qingniao

import "context"

type wechatDeliver struct {
	Title   string `validate:"required"`
	Content string `validate:"required"`

	executor wechatExecutor
}

func newWechatDeliver(title string, content string, executor wechatExecutor) *wechatDeliver {
	return &wechatDeliver{
		Title:   title,
		Content: content,

		executor: executor,
	}
}

func (wd *wechatDeliver) Send(ctx context.Context) (string, Status, error) {
	return wd.executor.send(ctx, wd)
}
