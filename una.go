package una

import (
	`context`

	`github.com/go-resty/resty/v2`
)

// Una 通知接口
type Una interface {
	// Send 发送消息
	Send(ctx context.Context, content string, opts ...option) (id string, err error)
}

// New 创建适配器
func New(resty *resty.Request) Una {
	return &unaTemplate{
		email:       NewEmail(),
		chuangcache: NewChuangcache(resty),
	}
}
