package una

import (
	`context`

	`github.com/go-resty/resty/v2`
	`github.com/storezhang/validatorx`
)

// Una 通知接口
type Una interface {
	// Send 发送消息
	Send(ctx context.Context, content string, opts ...option) (id string, err error)
}

// New 创建适配器
func New(config Config, resty *resty.Request) (una Una, err error) {
	if err = validatorx.Validate(config); nil != err {
		return
	}

	switch config.Type {
	case TypeChuangcache:
		una, err = NewChuangcacheSms(config.ChuangcacheSms, resty)
	}

	return
}
