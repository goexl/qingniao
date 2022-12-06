package qingniao

// Wechat 微信公众号推送
type Wechat struct {
	executor wechatExecutor
}

func newWechat(executor wechatExecutor) *Wechat {
	return &Wechat{
		executor: executor,
	}
}

func (w *Wechat) Deliver(title string, content string) *wechatDeliver {
	return newWechatDeliver(title, content, w.executor)
}
